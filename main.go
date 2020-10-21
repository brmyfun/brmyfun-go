package main

import (
	"fmt"
	"log"

	"github.com/brmyfun/brmy-go/config"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start brmyfun...")
	fmt.Println(config.Conf.Server)

	r := gin.Default()

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
