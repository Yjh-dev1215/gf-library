package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type IBooks interface {
	Add(ctx context.Context, in model.BooksAddInput) (res *model.BooksAddOutput, err error)
	Get(ctx context.Context, ISBN string) *entity.Books
	Update(ctx context.Context, in model.BooksUpdateInput) (res *model.BooksUpdateOutput, err error)
	Delete(ctx context.Context, Id int) (res *model.BooksDeleteOutput, err error)
}

var localBooks IBooks

func Books() IBooks {
	if localBooks == nil {
		panic("implement not found for interface IBooks, forgot register?")
	}
	return localBooks
}
func RegisterBooks(i IBooks) {
	localBooks = i
}
