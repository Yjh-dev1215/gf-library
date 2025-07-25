package books

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/books/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) AddBooks(ctx context.Context, req *v1.AddBooksReq) (res *v1.AddBooksRes, err error) {
	output, err := service.Books().Add(ctx, model.BooksAddInput{
		Title:    req.Title,
		Author:   req.Author,
		ISBN:     req.ISBN,
		Category: req.Category,
		Stock:    req.Stock,
		Total:    req.Total,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.AddBooksRes{
		Message: output.Message,
	}
	return
}
