package v1

import "github.com/gogf/gf/v2/frame/g"

type SignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Email     string `v:"required|email"` //注册必须包含email
	Nickname  string
	Code      string `v:"required|size:4|regex:^\\d{4}$"` // 邮箱验证码必须是4位数字
}
type SignUpRes struct{}
