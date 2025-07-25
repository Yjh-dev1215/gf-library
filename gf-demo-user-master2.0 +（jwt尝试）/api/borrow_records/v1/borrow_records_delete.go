package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteRecordsReq struct {
	g.Meta `path:"/borrow_records/{id}" method:"delete" tags:"BorrowRecords" summary:"借阅记录删除"`
	Id     int `json:"id" in:"path" v:"required#借阅记录ID不能为空"`
}
type DeleteRecordsRes struct {
}
