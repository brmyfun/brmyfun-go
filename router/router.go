package router

import (
	"log"

	"github.com/brmyfun/brmy-go/handler"

	"github.com/brmyfun/brmy-go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() {
	r := gin.Default()
	// 跨域配置
	r.Use(cors.Default())
	// 静态文件
	r.Use(static.Serve("/", static.LocalFile("./static", false)))
	// 加载鉴权中间件
	authMiddleware, err := middleware.InitAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// 不需要鉴权的接口
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/email/code", handler.EmailVcHandler)

	// 需要鉴权的接口
	api := r.Group("/api")
	api.GET("/refresh_token", authMiddleware.RefreshHandler)
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/hello", handler.HelloHandler)
	}

	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
