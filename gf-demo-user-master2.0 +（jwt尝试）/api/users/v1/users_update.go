package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta    `path:"/users/{id}" method:"put" tags:"UsersService" summary:"Update a user"`
	Id        int64  `json:"id"         v:"required|min:1" dc:"用户ID"`
	Name      string `json:"name"       v:"required|length:2,50" dc:"用户名"`
	Email     string `json:"email"      v:"required|email" dc:"邮箱"`
	StudentId string `json:"student_id" v:"required|length:6,16" dc:"学号"`
}
type UpdateRes struct{}
