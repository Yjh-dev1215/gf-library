package books

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type sBooks struct{}

func init() {
	service.RegisterBooks(New())
}

func New() *sBooks {
	return &sBooks{}
}

// AddBooks 添加书籍
func (s *sBooks) Add(ctx context.Context, in model.BooksAddInput) (res *model.BooksAddOutput, err error) {
	_, err = dao.Books.Ctx(ctx).Data(do.Books{
		Title:    in.Title,
		Author:   in.Author,
		Isbn:     in.ISBN,
		Category: in.Category,
		Stock:    in.Stock,
		Total:    in.Total,
	}).Insert()
	if err != nil {
		return nil, err
	}
	res = &model.BooksAddOutput{
		Message: "插入成功",
	}
	return
}

// 根据ISBN获取书籍
func (s *sBooks) Get(ctx context.Context, ISBN string) *entity.Books {
	var book *entity.Books
	err := dao.Books.Ctx(ctx).Where(do.Books{Isbn: ISBN}).Scan(&book)
	if err != nil {
		return nil
	}
	return book
}

// 更新书籍
func (s *sBooks) Update(ctx context.Context, in model.BooksUpdateInput) (res *model.BooksUpdateOutput, err error) {
	_, err = dao.Books.Ctx(ctx).Data(do.Books{
		Title:    in.Title,
		Author:   in.Author,
		Isbn:     in.ISBN,
		Category: in.Category,
		Stock:    in.Stock,
		Total:    in.Total,
	}).Where(do.Books{Id: in.Id}).Update()
	if err != nil {
		return nil, err
	}
	res = &model.BooksUpdateOutput{
		Message: "更新成功",
	}
	return
}

// 删除书籍
func (s *sBooks) Delete(ctx context.Context, Id int) (res *model.BooksDeleteOutput, err error) {
	_, err = dao.Books.Ctx(ctx).Where(do.Books{Id: Id}).Delete()
	if err != nil {
		return nil, err
	}
	res = &model.BooksDeleteOutput{
		Message: "删除成功",
	}
	return
}
