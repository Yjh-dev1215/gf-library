package borrowrecords

import borrowrecords "github.com/gogf/gf-demo-user/v2/api/borrow_records"

type ControllerV1 struct{}

func NewV1() borrowrecords.IBorrowRecordsV1 {
	return &ControllerV1{}
}
