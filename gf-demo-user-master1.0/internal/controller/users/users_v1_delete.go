package users

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/users/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Users().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
