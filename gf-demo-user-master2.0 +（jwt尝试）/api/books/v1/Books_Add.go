package v1

import "github.com/gogf/gf/v2/frame/g"

type AddBooksReq struct {
	g.Meta   `path:"/books" method:"post" tags:"BooksService" summary:"Add books"`
	Title    string `json:"title" v:"required#请输入书籍名称"`
	Author   string `json:"author" v:"required#请输入书籍作者"`
	ISBN     string `json:"isbn" v:"required#请输入书籍ISBN"`
	Category string `json:"category" v:"required#请输入书籍分类"`
	Stock    int    `json:"stock" v:"required#请输入书籍库存"`
	Total    int    `json:"total" v:"required#请输入书籍总量"`
}

type AddBooksRes struct {
	Message string `json:"message"` //返回插入成功
}
