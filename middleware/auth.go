package middleware

import (
	"errors"
	"log"
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
	Cookie  string `json:"cookie"`
	Expired string `json:"expired"`
}

// InitAuthMiddleware 初始化认证/鉴权配置
func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	// 获取配置文件中的 认证/鉴权 配置
	auth := config.Conf.Auth
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       auth.Realm,
		Key:         []byte(auth.Key),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: auth.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					auth.IdentityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// 认证
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
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// 鉴权
			return true
		},
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			cookie, err := c.Cookie("token")
			if err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, handler.Ok("登录成功", Token{Token: token, Cookie: cookie, Expired: t.Format(time.RFC3339)}))
		},
		RefreshResponse: func(c *gin.Context, code int, token string, t time.Time) {
			cookie, err := c.Cookie("token")
			if err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, handler.Ok("刷新成功", Token{Token: token, Cookie: cookie, Expired: t.Format(time.RFC3339)}))
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, handler.Ok("退出成功", nil))
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusOK, handler.Err(message))
		},
		TokenLookup:    "header: Authorization, query: token, cookie: token",
		TokenHeadName:  "Bearer",
		SendCookie:     true,
		SecureCookie:   false, // non HTTPS dev environments
		CookieHTTPOnly: true,  // JS can't modify
		CookieName:     "token",
		CookieDomain:   auth.CookieDomain,
		TimeFunc:       time.Now,
	})
}
