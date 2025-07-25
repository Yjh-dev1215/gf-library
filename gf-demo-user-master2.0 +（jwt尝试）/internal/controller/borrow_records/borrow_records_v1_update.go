package borrowrecords

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/borrow_records/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) UpdateRecords(ctx context.Context, req *v1.UpdateRecordsReq) (res *v1.UpdateRecordsRes, err error) {
	_, err = service.BorrowRecords().Update(ctx, model.UpdateRecordsInput{
		Id:         req.Id,
		UserId:     req.UserId,
		BookId:     req.BookId,
		BorrowDate: req.BorrowDate,
		ReturnDate: req.ReturnDate,
		DueDate:    req.DueDate,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return
}
