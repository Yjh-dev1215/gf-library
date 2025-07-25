package v1

import "github.com/gogf/gf/v2/frame/g"

type SendEmailCodeReq struct {
	g.Meta `path:"/user/email-code" method:"post" tags:"UserService" summary:"Send email verification code"`
	Email  string `v:"required|email#请输入正确的邮箱地址"` // 发送验证码的邮箱地址
}
type SendEmailCodeRes struct {
}
