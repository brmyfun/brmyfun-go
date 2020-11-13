package handler

import (
	"github.com/gin-gonic/gin"
)

// Ok 操作成功
func Ok(message string, data interface{}) map[string]interface{} {
	return gin.H{
		"code":     1,
		"messsage": message,
		"data":     data,
	}
}

// Err 操作失败
func Err(message string) map[string]interface{} {
	return gin.H{
		"code":    0,
		"message": message,
	}
}
