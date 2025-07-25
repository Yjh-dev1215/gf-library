package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddRecordsReq struct {
	g.Meta     `path:"/borrow_records" method:"post" tags:"BorrowRecords" summary:"借阅记录添加"`
	UserId     int    `json:"user_id" v:"required#用户ID不能为空"`
	BookId     int    `json:"book_id" v:"required#书籍ID不能为空"`
	BorrowDate string `json:"borrow_date" v:"required#借阅日期不能为空"`
	ReturnDate string `json:"return_date" v:"required#归还日期不能为空"`
	DueDate    string `json:"due_date" v:"required#到期日期不能为空"`
}

type AddRecordsRes struct {
}
