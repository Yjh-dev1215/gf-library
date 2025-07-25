package model

type UserCreateInput struct {
	Passport string
	Email    string
	Password string
	Nickname string
	Code     string
}

type UserSignInInput struct {
	/*
		Passport string
		Email    string
	*/
	Identifier string // 用户输入的内容（用户名 or 邮箱）
	Password   string
}

type UserCheckCodeInput struct {
	Email string
	Code  string
}

type UserChangePasswordInput struct {
	Email    string
	Code     string
	Password string
}
