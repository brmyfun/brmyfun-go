package handler

import (
	"github.com/gin-gonic/gin"
)

// Ok 操作成功
func Ok(message string, data interface{}) map[string]interface{} {
	return gin.H{
		"code":    1,
		"message": message,
		"data":    data,
	}
}

// Fail 操作失败
func Fail(message string, data interface{}) map[string]interface{} {
	return gin.H{
		"code":    0,
		"message": message,
		"data":    data,
	}
}

// Err 操作出错s
func Err(message string) map[string]interface{} {
	return gin.H{
		"code":    0,
		"message": message,
	}
}
