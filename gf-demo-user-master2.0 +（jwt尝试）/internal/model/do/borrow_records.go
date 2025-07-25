// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BorrowRecords is the golang structure of table borrow_records for DAO operations like Where/Data.
type BorrowRecords struct {
	g.Meta     `orm:"table:borrow_records, do:true"`
	Id         interface{} //
	UserId     interface{} //
	BookId     interface{} //
	BorrowDate *gtime.Time //
	ReturnDate *gtime.Time //
	DueDate    *gtime.Time //
	Status     interface{} //
}
