package handler

import (
	"net/http"

	"github.com/brmyfun/brmy-go/spider"
	"github.com/gin-gonic/gin"
)

// SpiderHandler 爬虫测试
func SpiderHandler(c *gin.Context) {
	spider.ZhihuHotRankV1()
	c.JSON(http.StatusOK, Ok("操作成功", nil))
	return
}
