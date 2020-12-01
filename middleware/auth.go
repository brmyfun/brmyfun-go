package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/brmyfun/brmy-go/config"
	"github.com/brmyfun/brmy-go/form"
	"github.com/brmyfun/brmy-go/handler"
	"github.com/brmyfun/brmy-go/model"
)

// Token 封装令牌
type Token struct {
	Token   string `json:"token"`
	Expired string `json:"expired"`
}

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
				return nil, errors.New("用户名或密码不能为空")
			}
			username := loginForm.Username
			password := loginForm.Password
			ok, err := handler.LoginCheck(username, password)
			if ok {
				return &model.User{
					Username: username,
				}, nil
			}
			return nil, err
		},
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			c.JSON(http.StatusOK, handler.Ok("登录成功", Token{Token: token, Expired: t.Format(time.RFC3339)}))
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusOK, handler.Err(message))
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
