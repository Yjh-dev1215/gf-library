package users

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/users/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	err = service.Users().Add(ctx, model.UsersAddInput{
		Name:      req.Name,
		Email:     req.Email,
		StudentId: req.StudentId,
	})

	return
}
