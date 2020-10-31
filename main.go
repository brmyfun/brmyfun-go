package main

import (
	"fmt"
	"log"

	"github.com/brmyfun/brmy-go/config"

	"github.com/brmyfun/brmy-go/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	identityKey := config.Conf.Auth.IdentityKey
	c.JSON(200, gin.H{
		"username": claims[identityKey],
		"text":     "Hello World.",
	})
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	authMiddleware, err := middleware.InitAuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "NOT_FOUND", "message": "404 not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	r.GET("/", func(c *gin.Context) {
		fmt.Println("请求路径:", c.FullPath())
		c.JSON(200, gin.H{
			"code":    1,
			"message": "操作成功",
			"data":    "index",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
