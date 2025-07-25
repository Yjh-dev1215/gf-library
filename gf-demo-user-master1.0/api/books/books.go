package books

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/books/v1"
)

type IBooksV1 interface {
	AddBooks(ctx context.Context, req *v1.AddBooksReq) (res *v1.AddBooksRes, err error)                  //添加书籍
	DeleteBooks(ctx context.Context, req *v1.DeleteBooksReq) (res *v1.DeleteBooksRes, err error)         //删除书籍
	UpdateBooks(ctx context.Context, req *v1.UpdateBooksReq) (res *v1.UpdateBooksRes, err error)         //更新书籍
	GetBookByISBN(ctx context.Context, req *v1.GetBooksByISBNReq) (res *v1.GetBooksByISBNRes, err error) //根据ISBN获取书籍
}
