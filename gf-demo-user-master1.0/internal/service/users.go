package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type IUsers interface {
	Add(ctx context.Context, in model.UsersAddInput) (err error)
	Delete(ctx context.Context, Id int64) (err error)
	Update(ctx context.Context, in model.UsersUpdateInput) (err error)
	GetUserById(ctx context.Context, Id int64) *entity.Users
}

var localUsers IUsers

func Users() IUsers {
	if localUsers == nil {
		panic("implement not found for interface IUsers, forgot register?")
	}
	return localUsers
}

func RegisterUsers(i IUsers) {
	localUsers = i
}
