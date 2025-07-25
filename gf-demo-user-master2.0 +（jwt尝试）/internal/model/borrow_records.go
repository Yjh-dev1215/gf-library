package model

type AddRecordsInput struct {
	UserId     int    `json:"user_id" v:"required#用户ID不能为空"`
	BookId     int    `json:"book_id" v:"required#书籍ID不能为空"`
	BorrowDate string `json:"borrow_date" v:"required#借阅日期不能为空"`
	ReturnDate string `json:"return_date" v:"required#归还日期不能为空"`
	DueDate    string `json:"due_date" v:"required#到期日期不能为空"`
}

type UpdateRecordsInput struct {
	Id         int    `json:"id" v:"required#记录ID不能为空"`
	UserId     int    `json:"user_id" v:"required#用户ID不能为空"`
	BookId     int    `json:"book_id" v:"required#书籍ID不能为空"`
	BorrowDate string `json:"borrow_date" v:"required#借阅日期不能为空"`
	ReturnDate string `json:"return_date" v:"required#归还日期不能为空"`
	DueDate    string `json:"due_date" v:"required#到期日期不能为空"`
	Status     string `json:"status" v:"required#状态不能为空"` //enum('borrowed', 'returned', 'overdue')
}
