package user

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/errors/gerror"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	if req.Identifier == "" {
		return nil, gerror.New("账号或邮箱必须填写一个")
	}
	_, err = service.User().SignIn(ctx, model.UserSignInInput{
		/*
			Passport: req.Passport,
			Email:    req.Email,
		*/
		Identifier: req.Identifier,
		Password:   req.Password,
	})
	return
}
