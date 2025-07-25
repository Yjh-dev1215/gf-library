package v1

import "github.com/gogf/gf/v2/frame/g"

type ResetPasswordReq struct {
	g.Meta    `path:"/user/reset-password" method:"post" tags:"UserService" summary:"Reset user password"`
	Email     string `v:"required|email"`                     // Email is required for password reset
	Code      string `v:"required|length:4,4"`                // 邮箱验证码，必须是4位
	Password  string `v:"required|length:6,16"`               // New password
	Password2 string `v:"required|length:6,16|same:Password"` // Confirm new password
}

type ResetPasswordRes struct {
}
