package books

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/books/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) GetBookByISBN(ctx context.Context, req *v1.GetBooksByISBNReq) (res *v1.GetBooksByISBNRes, err error) {
	res = &v1.GetBooksByISBNRes{
		Books: service.Books().Get(ctx, req.ISBN),
	}
	return
}
