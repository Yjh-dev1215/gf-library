package user

import (
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jordan-wright/email"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

var redisClient = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
})

// Create creates user account.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	var redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	checkInput := model.UserCheckCodeInput{
		Email: in.Email,
		Code:  in.Code,
	}
	if err := s.IsCodeRight(ctx, checkInput); err != nil {
		return err
	}

	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)
	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 在写入数据库前加密密码
		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return gerror.New("密码加密失败")
		}
		in.Password = string(hashedPwd)
		//写入
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			Email:    in.Email, //新增 必须
			Password: in.Password,
			Nickname: in.Nickname,
		}).Insert()
		if err != nil {
			return err
		}
		// 注册成功后删除验证码
		redisKey := "register:email:" + in.Email
		_, _ = redisClient.Del(ctx, redisKey).Result()
		return nil
	})

}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// 登录统一接口 Identifier
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.User
	// 正则匹配判断是否为邮箱
	isEmail := gregex.IsMatchString(`^[\w.-]+@[\w.-]+\.[a-zA-Z]{2,}$`, in.Identifier)
	if isEmail {
		// 邮箱登录
		err = dao.User.Ctx(ctx).Where(do.User{
			Email: in.Identifier,
		}).Scan(&user)
	} else {
		// 用户名登录
		err = dao.User.Ctx(ctx).Where(do.User{
			Passport: in.Identifier,
		}).Scan(&user)
	}
	if err != nil {
		return
	}
	if user == nil {
		return gerror.New(`账号或密码错误`)
	}
	// 用 bcrypt 校验密码
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)) != nil {
		return gerror.New(`账号或密码错误`)
	}
	//
	if err = service.Session().SetUser(ctx, user); err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx) //注销登录 → 删除当前用户 Session
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsCodeRight checks if the provided email code is correct.
func (s *sUser) IsCodeRight(ctx context.Context, in model.UserCheckCodeInput) error {
	var redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	redisKey := "register:email:" + in.Email
	fmt.Printf("验证码检查: 邮箱=%s, 提交的验证码=%s\n", in.Email, in.Code)

	codeInRedis, err := redisClient.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		fmt.Printf("验证码已过期或不存在: %s\n", redisKey)
		return gerror.New("验证码已过期或不存在")
	}
	if err != nil {
		fmt.Printf("获取验证码错误: %v\n", err)
		return err
	}

	fmt.Printf("Redis中的验证码: %s (类型: %T)\n", codeInRedis, codeInRedis)

	// 确保验证码格式一致，去除可能的空白字符
	trimmedCodeInRedis := strings.TrimSpace(codeInRedis)
	trimmedInputCode := strings.TrimSpace(in.Code)

	fmt.Printf("比较: Redis验证码=%s, 提交的验证码=%s\n", trimmedCodeInRedis, trimmedInputCode)

	if trimmedCodeInRedis != trimmedInputCode {
		fmt.Printf("验证码不匹配: Redis=%s, 提交=%s\n", trimmedCodeInRedis, trimmedInputCode)
		return gerror.New("验证码错误")
	}
	/*
		fmt.Printf("验证码匹配成功，准备删除Redis中的验证码\n")
		// 校验成功后，删除验证码
		_, err = redisClient.Del(ctx, redisKey).Result()
		if err != nil {
			fmt.Printf("删除验证码失败: %v\n", err)
		}
	*/
	return nil

}

// GetProfile retrieves and returns current user info in session.
func (s *sUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx) //从 Session 中获取当前登录用户信息 → 用于展示个人资料页
}

func (s *sUser) SendEmailCode(ctx context.Context, emailAddr string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendEmailCode panic:", r)
		}
	}()
	fmt.Println("Controller SendEmailCode 被调用")
	// 1. 生成4位随机验证码
	// 使用time.Now().UnixNano()作为种子创建新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Intn(9000) + 1000 // 1000~9999

	// 2. 发送邮件
	e := email.NewEmail()
	e.From = "Library_system <yjh201004030080@qq.com>"
	e.To = []string{emailAddr}
	e.Subject = "Library_system|邮箱验证码"
	e.Text = []byte(fmt.Sprintf("您的验证码是：%d,有效期5分钟。", code))
	err = e.Send("smtp.qq.com:587", smtp.PlainAuth("", "yjh201004030080@qq.com", "acpjtbdbcsztfibi", "smtp.qq.com"))
	/*
				if err != nil {
					return err
				} //你的邮件发送 e.Send(...) 实际上返回了一个非 nil 的 error（即使你能收到邮件，go-redis/jordan-wright/email 这类库有时会返回 warning 或非致命错误）。
		只要 err != nil，就会 return err，后面的 Redis 写入代码根本不会执行，所以你看不到任何 Redis 相关日志。
		你注释掉这段后，无论邮件发送是否报错，都会继续执行 Redis 写入，所以日志和功能都正常了。
	*/
	// 邮箱验证码功能
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	redisKey := "register:email:" + emailAddr
	fmt.Println("准备写入Redis")
	err = client.Set(context.Background(), redisKey, code, 5*time.Minute).Err()
	fmt.Println("Set执行完毕", err)
	if err != nil {
		fmt.Println("存储验证码失败:", err)
		return err
	}
	fmt.Println("存储验证码成功")
	return nil
}

func (s *sUser) ChangePassword(ctx context.Context, in model.UserChangePasswordInput) (err error) {
	fmt.Println("ChangePassword:", in.Email, in.Password)
	// 1. 校验验证码

	checkInput := model.UserCheckCodeInput{
		Email: in.Email,
		Code:  in.Code,
	}

	if err := s.IsCodeRight(ctx, checkInput); err != nil {
		return err
	}

	// 2. 校验邮箱是否存在
	count, err := dao.User.Ctx(ctx).Where(do.User{Email: in.Email}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.Newf(`邮箱 "%s" 未注册`, in.Email)
	}
	// 3. 对新密码进行hash加密
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密失败:", err)
		return gerror.New("密码加密失败")
	}

	// 4. 更新密码
	fmt.Println("开始更新密码，邮箱:", in.Email, "新密码:", in.Password)
	_, err = dao.User.Ctx(ctx).Where(do.User{Email: in.Email}).Update(do.User{Password: string(hashedPwd)})
	if err != nil {
		fmt.Println("更新密码失败:", err)
		return err
	}

	// 5. 删除Redis中的验证码
	redisKey := "register:email:" + in.Email
	_, _ = redisClient.Del(ctx, redisKey).Result()
	return nil
}
