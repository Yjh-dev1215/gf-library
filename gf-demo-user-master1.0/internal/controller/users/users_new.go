package users

import "github.com/gogf/gf-demo-user/v2/api/users"

type ControllerV1 struct{}

func NewV1() users.IUsersV1 {
	return &ControllerV1{}
}
