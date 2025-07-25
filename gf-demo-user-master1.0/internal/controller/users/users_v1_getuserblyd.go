package users

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/users/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) GetUserById(ctx context.Context, req *v1.GetByIdReq) (res *v1.GetByIdRes, err error) {
	res = &v1.GetByIdRes{
		Users: service.Users().GetUserById(ctx, req.Id),
	}
	return
}
