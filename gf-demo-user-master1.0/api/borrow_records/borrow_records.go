package borrowrecords

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/borrow_records/v1"
)

type IBorrowRecordsV1 interface {
	AddRecords(ctx context.Context, req *v1.AddRecordsReq) (res *v1.AddRecordsRes, err error)
	GetRecords(ctx context.Context, req *v1.GetRecordsReq) (res *v1.GetRecordsRes, err error)
	DeleteRecords(ctx context.Context, req *v1.DeleteRecordsReq) (res *v1.DeleteRecordsRes, err error)
	UpdateRecords(ctx context.Context, req *v1.UpdateRecordsReq) (res *v1.UpdateRecordsRes, err error)
}
