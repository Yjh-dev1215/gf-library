package model

type UsersAddInput struct {
	Name      string `json:"name" v:"required"`
	Email     string `json:"email" v:"required"`
	StudentId string `json:"student_id" v:"required"`
}

type UsersUpdateInput struct {
	Id        int64  `json:"id" v:"required"`
	Name      string `json:"name" v:"required"`
	Email     string `json:"email" v:"required"`
	StudentId string `json:"student_id" v:"required"`
}
