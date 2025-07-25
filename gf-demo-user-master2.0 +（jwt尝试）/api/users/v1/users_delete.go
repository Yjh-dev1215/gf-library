package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteReq struct {
	g.Meta `path:"/users/{id}" method:"delete" tags:"usersservice" summary:"Delete a user"`
	Id     int64 `v:"required"` // 用户ID
}
type DeleteRes struct{}
