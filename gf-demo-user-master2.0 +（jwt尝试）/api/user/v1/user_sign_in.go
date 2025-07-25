package v1

import "github.com/gogf/gf/v2/frame/g"

type SignInReq struct {
	g.Meta `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	// Passport string `v:"required"`
	/*
		Passport string `v:"length:6,16"`
		Email    string `v:"email"` //登录时 账号或email可选
	*/
	Identifier string `v:"required|length:6,30|email#请输入账号或邮箱"` //登录时 账号或email可选
	Password   string `v:"required"`
}
type SignInRes struct{}
