package users

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/users/v1"
)

type IUsersV1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)                 //添加用户
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)        //删除用户
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)        //更新用户
	GetUserById(ctx context.Context, req *v1.GetByIdReq) (res *v1.GetByIdRes, err error) //根据ID获取用户
}
