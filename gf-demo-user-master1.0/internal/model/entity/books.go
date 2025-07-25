// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Books is the golang structure for table books.
type Books struct {
	Id       int    `json:"id"       orm:"id"       description:""`
	Title    string `json:"title"    orm:"title"    description:""`
	Author   string `json:"author"   orm:"author"   description:""`
	Isbn     string `json:"isbn"     orm:"isbn"     description:""`
	Category string `json:"category" orm:"category" description:""`
	Stock    int    `json:"stock"    orm:"stock"    description:""`
	Total    int    `json:"total"    orm:"total"    description:""`
}
