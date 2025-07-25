package borrowrecords

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/borrow_records/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) DeleteRecords(ctx context.Context, req *v1.DeleteRecordsReq) (res *v1.DeleteRecordsRes, err error) {
	_, err = service.BorrowRecords().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	/*if req.Id != ctx.Value("id_from_path") {
	      return nil, gerror.New("路径参数和body参数不一致")
	  }
	*/
	return
}
