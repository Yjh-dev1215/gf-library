package borrow_records

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/borrow_records/v1"
	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sBorrowRecords struct{}

func init() {
	service.RegisterBorrowRecords(New())
}

func New() *sBorrowRecords {
	return &sBorrowRecords{}
}

// 添加
func (s *sBorrowRecords) Add(ctx context.Context, in model.AddRecordsInput) (res *v1.AddRecordsRes, err error) {
	_, err = dao.BorrowRecords.Ctx(ctx).Data(do.BorrowRecords{
		UserId:     in.UserId,
		BookId:     in.BookId,
		BorrowDate: gtime.NewFromStr(in.BorrowDate),
		ReturnDate: gtime.NewFromStr(in.ReturnDate),
		DueDate:    gtime.NewFromStr(in.DueDate),
	}).Insert()
	if err != nil {
		return nil, err
	}
	res = &v1.AddRecordsRes{}
	return
}

// 修改
func (s *sBorrowRecords) Update(ctx context.Context, in model.UpdateRecordsInput) (res *v1.UpdateRecordsRes, err error) {
	_, err = dao.BorrowRecords.Ctx(ctx).Data(do.BorrowRecords{
		UserId:     in.UserId,
		BookId:     in.BookId,
		BorrowDate: gtime.NewFromStr(in.BorrowDate),
		ReturnDate: gtime.NewFromStr(in.ReturnDate),
		DueDate:    gtime.NewFromStr(in.DueDate),
		Status:     in.Status,
	}).Where(do.BorrowRecords{Id: in.Id}).Update()
	if err != nil {
		return nil, err
	}
	res = &v1.UpdateRecordsRes{}
	return
}

// 删除
func (s *sBorrowRecords) Delete(ctx context.Context, Id int) (res *v1.DeleteRecordsRes, err error) {
	_, err = dao.BorrowRecords.Ctx(ctx).Where(do.BorrowRecords{Id: Id}).Delete()
	if err != nil {
		return nil, err
	}
	res = &v1.DeleteRecordsRes{}
	return
}

// 查询
func (s *sBorrowRecords) Get(ctx context.Context, req *v1.GetRecordsReq) (records []*entity.BorrowRecords, err error) {
	m := dao.BorrowRecords.Ctx(ctx)
	if req.Id == nil && req.BookId == nil {
		return nil, gerror.New("请至少传入一个查询条件")
	}
	if req.Id != nil {
		m = m.Where("id", *req.Id)
	}
	if req.BookId != nil {
		m = m.Where("book_id", *req.BookId)
	}
	err = m.Scan(&records)
	return
}
