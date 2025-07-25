package user

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) CheckEmailCode(ctx context.Context, req *v1.CheckEmailCodeReq) (res *v1.CheckEmailCodeRes, err error) {
	err = service.User().IsCodeRight(ctx, model.UserCheckCodeInput{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		return nil, gerror.Newf(`验证码错误`)
	}
	return
}
