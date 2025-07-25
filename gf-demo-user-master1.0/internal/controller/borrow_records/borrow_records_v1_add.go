package borrowrecords

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/borrow_records/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) AddRecords(ctx context.Context, req *v1.AddRecordsReq) (res *v1.AddRecordsRes, err error) {
	_, err = service.BorrowRecords().Add(ctx, model.AddRecordsInput{
		UserId:     req.UserId,
		BookId:     req.BookId,
		BorrowDate: req.BorrowDate,
		ReturnDate: req.ReturnDate,
		DueDate:    req.DueDate,
	})
	if err != nil {
		return nil, err
	}
	return
}
