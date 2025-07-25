package user

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) SendEmailCode(ctx context.Context, req *v1.SendEmailCodeReq) (res *v1.SendEmailCodeRes, err error) {
	err = service.User().SendEmailCode(ctx, req.Email)
	return
}
