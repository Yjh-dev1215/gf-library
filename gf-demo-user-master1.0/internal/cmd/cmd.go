package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/gogf/gf-demo-user/v2/internal/consts"
	"github.com/gogf/gf-demo-user/v2/internal/controller/books"
	borrowrecords "github.com/gogf/gf-demo-user/v2/internal/controller/borrow_records"
	"github.com/gogf/gf-demo-user/v2/internal/controller/user"
	"github.com/gogf/gf-demo-user/v2/internal/controller/users"
	_ "github.com/gogf/gf-demo-user/v2/internal/logic/books"
	_ "github.com/gogf/gf-demo-user/v2/internal/logic/borrow_records"
	_ "github.com/gogf/gf-demo-user/v2/internal/logic/users"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of simple goframe demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				// Register route handlers.
				var (
					userCtrl          = user.NewV1()
					usersCtrl         = users.NewV1()
					booksCtrl         = books.NewV1()
					borrowrecordsCtrl = borrowrecords.NewV1()
				)
				group.Bind(
					userCtrl,
					usersCtrl,
					booksCtrl,
					borrowrecordsCtrl,
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/user/profile": userCtrl.Profile,
					})
				})
			})
			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			// Just run the server.
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "GoFrame",
			URL:  "https://goframe.org",
		},
	}
}
