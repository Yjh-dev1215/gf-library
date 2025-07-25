package v1

import "github.com/gogf/gf/v2/frame/g"

type AddReq struct {
	g.Meta    `path:"/users" method:"post" tags:"UsersService" summary:"Add a new user"`
	Name      string `v:"required|length:1,16"` // 用户名
	Email     string `v:"required|email"`       // 邮箱
	StudentId string `v:"required|length:6,16"` // 学号
}
type AddRes struct{}
