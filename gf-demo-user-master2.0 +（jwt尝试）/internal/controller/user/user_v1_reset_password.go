package user

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error) {
	err = service.User().ChangePassword(ctx, model.UserChangePasswordInput{
		Email:    req.Email,
		Code:     req.Code,
		Password: req.Password,
	})
	return
}
