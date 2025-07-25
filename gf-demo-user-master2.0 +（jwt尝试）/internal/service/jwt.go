package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
)

type IJwt interface {
	GenerateToken(ctx context.Context, userId uint) (string, error)
	ParseToken(ctx context.Context, tokenString string) (*jwt.MapClaims, error)
}

type jwtService struct{}

var (
	jwtSecret = []byte("your-secret-key") // 应该从配置中读取
)

func Jwt() IJwt {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(ctx context.Context, userId uint) (string, error) {
	// 设置token过期时间
	expireTime := time.Now().Add(24 * time.Hour) // 24小时过期

	// 创建claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     expireTime.Unix(),
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (s *jwtService) ParseToken(ctx context.Context, tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid token")
}
