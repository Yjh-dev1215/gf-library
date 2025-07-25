// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Books is the golang structure of table books for DAO operations like Where/Data.
type Books struct {
	g.Meta   `orm:"table:books, do:true"`
	Id       interface{} //
	Title    interface{} //
	Author   interface{} //
	Isbn     interface{} //
	Category interface{} //
	Stock    interface{} //
	Total    interface{} //
}
