package v1

import (
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetRecordsReq struct {
	g.Meta `path:"/borrow_records" method:"get" tags:"BorrowRecords" summary:"借阅记录查询"`
	Id     *int `json:"id,omitempty" dc:"借阅记录ID"`    // 可选，借阅记录ID 查询单条记录
	BookId *int `json:"book_id,omitempty" dc:"书籍ID"` // 可选，书籍ID 查询多条记录
}

/*
	type GetRecordsRes struct {
		*entity.BorrowRecords
	}
*/
type GetRecordsRes struct {
	List []*entity.BorrowRecords `json:"list"`
}
