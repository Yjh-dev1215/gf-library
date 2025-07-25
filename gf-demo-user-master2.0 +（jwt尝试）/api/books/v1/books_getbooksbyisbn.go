package v1

import (
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetBooksByISBNReq struct {
	g.Meta `path:"/books/{ISBN}" tags:"Books" method:"get" summary:"根据ISBN获取书籍"`
	ISBN   string `json:"isbn" v:"required|length:10,13" dc:"ISBN"` // ISBN
}

type GetBooksByISBNRes struct {
	*entity.Books
}
