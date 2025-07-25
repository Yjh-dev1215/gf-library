package books

import "github.com/gogf/gf-demo-user/v2/api/books"

type ControllerV1 struct{}

func NewV1() books.IBooksV1 {
	return &ControllerV1{}
}
