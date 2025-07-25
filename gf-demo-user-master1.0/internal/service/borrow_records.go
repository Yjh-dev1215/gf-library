package service

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/borrow_records/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type IBorrowRecords interface {
	Add(ctx context.Context, in model.AddRecordsInput) (res *v1.AddRecordsRes, err error)
	Update(ctx context.Context, in model.UpdateRecordsInput) (res *v1.UpdateRecordsRes, err error)
	Delete(ctx context.Context, Id int) (res *v1.DeleteRecordsRes, err error)
	Get(ctx context.Context, req *v1.GetRecordsReq) (records []*entity.BorrowRecords, err error)
}

var localBorrowRecords IBorrowRecords //接口实现和获取

func BorrowRecords() IBorrowRecords {
	if localBorrowRecords == nil {
		panic("implement not found for interface IBorrowRecords")
	}
	return localBorrowRecords
}

func RegisterBorrowRecords(i IBorrowRecords) {
	localBorrowRecords = i
}
