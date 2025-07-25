package v1

import "github.com/gogf/gf/v2/frame/g"

type CheckEmailCodeReq struct {
	g.Meta `path:"/user/check-email-code" method:"post" tags:"UserService" summary:"Check email code"`
	Email  string `v:"required|email" json:"email" dc:"Email address"`
	Code   string `v:"required|size:4|regex:^\\d{4}$" json:"code" dc:"Email verification code must be 4 digits"`
}

type CheckEmailCodeRes struct {
}
