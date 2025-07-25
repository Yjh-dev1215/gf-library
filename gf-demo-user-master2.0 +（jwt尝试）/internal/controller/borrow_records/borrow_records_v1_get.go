package borrowrecords

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/borrow_records/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

/*
	func (c *ControllerV1) GetRecords(ctx context.Context, req *v1.GetRecordsReq) (res *v1.GetRecordsRes, err error) {
		record, err := service.BorrowRecords().Get(ctx, req)
		if err != nil {
			return nil, err
		}
		res = &v1.GetRecordsRes{
			BorrowRecords: record,
		}
		return
	}
*/
func (c *ControllerV1) GetRecords(ctx context.Context, req *v1.GetRecordsReq) (res *v1.GetRecordsRes, err error) {
	records, err := service.BorrowRecords().Get(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.GetRecordsRes{
		List: records,
	}
	return
}
