package middleware

import (
	"time"

	"github.com/brmyfun/brmy-go/handler"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/brmyfun/brmy-go/config"
	"github.com/gin-gonic/gin"
)

// InitAuthMiddleware 初始化鉴权配置
func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	auth := config.Conf.Auth
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:         auth.Realm,
		Key:           []byte(auth.Key),
		IdentityKey:   auth.IdentityKey,
		Authenticator: handler.LoginHandler(c * gin.Context),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
