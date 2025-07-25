package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
}

// Auth validates the request using JWT token.
func (s *sMiddleware) Auth(r *ghttp.Request) {
	// Get token from Authorization header
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	// Verify JWT token
	claims, err := service.Jwt().ParseToken(r.Context(), token)
	if err != nil {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	// Get user ID from claims
	userId, ok := (*claims)["user_id"].(float64)
	if !ok {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	// Set user info to context
	service.BizCtx().SetUser(r.Context(), &model.ContextUser{
		Id: uint(userId),
	})

	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
