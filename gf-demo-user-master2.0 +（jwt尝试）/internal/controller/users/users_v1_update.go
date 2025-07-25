package users

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/users/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.Users().Update(ctx, model.UsersUpdateInput{
		Id:        req.Id,
		Name:      req.Name,
		Email:     req.Email,
		StudentId: req.StudentId,
	})
	if err != nil {
		return nil, err
	}
	return
}
