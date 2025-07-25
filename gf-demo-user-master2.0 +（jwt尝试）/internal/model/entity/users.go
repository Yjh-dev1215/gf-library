// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	Name         string      `json:"name"         orm:"name"          description:""`
	Email        string      `json:"email"        orm:"email"         description:""`
	StudentId    string      `json:"studentId"    orm:"student_id"    description:""`
	RegisterTime *gtime.Time `json:"registerTime" orm:"register_time" description:""`
}
