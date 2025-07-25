// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BorrowRecordsDao is the data access object for the table borrow_records.
type BorrowRecordsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  BorrowRecordsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// BorrowRecordsColumns defines and stores column names for the table borrow_records.
type BorrowRecordsColumns struct {
	Id         string //
	UserId     string //
	BookId     string //
	BorrowDate string //
	ReturnDate string //
	DueDate    string //
	Status     string //
}

// borrowRecordsColumns holds the columns for the table borrow_records.
var borrowRecordsColumns = BorrowRecordsColumns{
	Id:         "id",
	UserId:     "user_id",
	BookId:     "book_id",
	BorrowDate: "borrow_date",
	ReturnDate: "return_date",
	DueDate:    "due_date",
	Status:     "status",
}

// NewBorrowRecordsDao creates and returns a new DAO object for table data access.
func NewBorrowRecordsDao(handlers ...gdb.ModelHandler) *BorrowRecordsDao {
	return &BorrowRecordsDao{
		group:    "default",
		table:    "borrow_records",
		columns:  borrowRecordsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BorrowRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BorrowRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BorrowRecordsDao) Columns() BorrowRecordsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BorrowRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BorrowRecordsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BorrowRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
