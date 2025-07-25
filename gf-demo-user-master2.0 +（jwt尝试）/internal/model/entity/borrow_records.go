// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BorrowRecords is the golang structure for table borrow_records.
type BorrowRecords struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	UserId     int         `json:"userId"     orm:"user_id"     description:""`
	BookId     int         `json:"bookId"     orm:"book_id"     description:""`
	BorrowDate *gtime.Time `json:"borrowDate" orm:"borrow_date" description:""`
	ReturnDate *gtime.Time `json:"returnDate" orm:"return_date" description:""`
	DueDate    *gtime.Time `json:"dueDate"    orm:"due_date"    description:""`
	Status     string      `json:"status"     orm:"status"      description:""`
}
