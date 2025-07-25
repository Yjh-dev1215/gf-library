package books

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/books/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) DeleteBooks(ctx context.Context, req *v1.DeleteBooksReq) (res *v1.DeleteBooksRes, err error) {
	output, err := service.Books().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.DeleteBooksRes{
		Message: output.Message,
	}
	return
}
