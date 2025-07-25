package books

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/books/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) UpdateBooks(ctx context.Context, req *v1.UpdateBooksReq) (res *v1.UpdateBooksRes, err error) {
	output, err := service.Books().Update(ctx, model.BooksUpdateInput{
		Id:       req.Id,
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
	res = &v1.UpdateBooksRes{
		Message: output.Message,
	}
	return
}
