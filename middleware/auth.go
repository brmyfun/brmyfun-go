package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/brmyfun/brmy-go/config"
	"github.com/brmyfun/brmy-go/form"
	"github.com/brmyfun/brmy-go/handler"
	"github.com/brmyfun/brmy-go/model"
)

// InitAuthMiddleware 初始化鉴权配置
func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	auth := config.Conf.Auth
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       auth.Realm,
		Key:         []byte(auth.Key),
		IdentityKey: auth.IdentityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginForm form.LoginDefault
			if err := c.ShouldBind(&loginForm); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginForm.Username
			password := loginForm.Password
			ok, err := handler.LoginPrecheck(username, password)
			if ok {
				return &model.User{
					Username: username,
				}, nil
			}
			return nil, err
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
