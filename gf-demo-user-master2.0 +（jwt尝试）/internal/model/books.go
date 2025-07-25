package model

type BooksAddInput struct {
	Title    string `json:"title" v:"required#请输入书籍名称"`
	Author   string `json:"author" v:"required#请输入书籍作者"`
	ISBN     string `json:"isbn" v:"required#请输入书籍ISBN"`
	Category string `json:"category" v:"required#请输入书籍分类"`
	Stock    int    `json:"stock" v:"required#请输入书籍库存"`
	Total    int    `json:"total" v:"required#请输入书籍总量"`
}

type BooksUpdateInput struct {
	Id       int    `json:"id" v:"required#请输入书籍ID"`
	Title    string `json:"title" v:"required#请输入书籍名称"`
	Author   string `json:"author" v:"required#请输入书籍作者"`
	ISBN     string `json:"isbn" v:"required#请输入书籍ISBN"`
	Category string `json:"category" v:"required#请输入书籍分类"`
	Stock    int    `json:"stock" v:"required#请输入书籍库存"`
	Total    int    `json:"total" v:"required#请输入书籍总量"`
}
type BooksAddOutput struct {
	Message string `json:"message"`
}
type BooksUpdateOutput struct {
	Message string `json:"message"`
}
type BooksDeleteOutput struct {
	Message string `json:"message"`
}
