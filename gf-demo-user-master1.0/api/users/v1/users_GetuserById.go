package v1

import (
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetByIdReq struct {
	g.Meta `path:"/users/{id}" method:"get" tags:"Users" summary:"根据 ID 获取用户详情"`
	Id     int64 `json:"id" v:"required|min:1"` // 用户 ID，必填且最小值为 1
}

// 根据ID获取用户响应体
type GetByIdRes struct {
	*entity.Users
}
