package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateBooksReq struct {
	g.Meta   `path:"/books/{id}" tags:"Books" method:"put" summary:"更新书籍"`
	Id       int    `json:"id" in:"path" v:"required|min:1" dc:"表ID"`        // 表ID
	Title    string `json:"title"       v:"required|length:2,50" dc:"书名"`    // 书名
	Author   string `json:"author"      v:"required|length:2,50" dc:"作者"`    // 作者
	ISBN     string `json:"isbn"        v:"required|length:10,13" dc:"ISBN"` // ISBN
	Category string `json:"category"    v:"required|length:2,50" dc:"分类"`    // 分类
	Stock    int    `json:"stock"       v:"required|integer" dc:"库存"`        // 库存
	Total    int    `json:"total"       v:"required|integer" dc:"总量"`        // 总量
}

type UpdateBooksRes struct {
	Message string `json:"message" dc:"更新成功"` // 更新成功
}
