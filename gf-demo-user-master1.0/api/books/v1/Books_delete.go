package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteBooksReq struct {
	g.Meta `path:"/books/{id}" method:"delete" tags:"BooksService" summary:"Delete books"`
	Id     int `v:"required"` // Id更高效
}

type DeleteBooksRes struct {
	Message string `json:"message"` //返回删除成功
}
