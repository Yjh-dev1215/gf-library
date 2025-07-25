package users

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type sUsers struct{}

func init() {
	service.RegisterUsers(New())
}
func New() *sUsers {
	return &sUsers{}
}

// 添加用户
func (s *sUsers) Add(ctx context.Context, in model.UsersAddInput) (err error) {
	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Name:      in.Name,
		Email:     in.Email,
		StudentId: in.StudentId,
	}).Insert()
	return
}

// 删除用户
func (s *sUsers) Delete(ctx context.Context, id int64) (err error) {
	_, err = dao.Users.Ctx(ctx).Where(do.Users{Id: id}).Delete()
	return
}

// 更新用户
func (s *sUsers) Update(ctx context.Context, in model.UsersUpdateInput) (err error) {
	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Name:      in.Name,
		Email:     in.Email,
		StudentId: in.StudentId,
	}).Where(do.Users{Id: in.Id}).Update()
	return
}

// 根据ID获取用户
func (s *sUsers) GetUserById(ctx context.Context, id int64) *entity.Users {
	var user *entity.Users
	_ = dao.Users.Ctx(ctx).Where(do.Users{Id: id}).Scan(&user)
	return user
}
