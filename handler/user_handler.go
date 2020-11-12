package handler

import (
	"net/http"
	"time"

	"github.com/brmyfun/brmy-go/config"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/brmyfun/brmy-go/model"
	"github.com/brmyfun/brmy-go/service"
	"github.com/go-redis/redis/v8"

	"github.com/brmyfun/brmy-go/form"
	"github.com/brmyfun/brmy-go/util"
	"github.com/gin-gonic/gin"
)

// VerificationCodeHandler 验证码处理器
func VerificationCodeHandler(c *gin.Context) {
	telephone := c.PostForm("telephone")
	if !util.VerifyTelephone(telephone) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "手机号格式有误",
			"data":    nil,
		})
	}
	// 先从缓存中获取验证码
	vc, err := service.RedisGet(telephone)
	if err == redis.Nil {
		// 缓存中没有则生成一个新的验证码
		vc = util.GenRandomNumCode(6)
		// 存入缓存
		_, err := service.RedisSet(telephone, vc, time.Minute)
		if err != nil {
			// 这个错误需要处理一下
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "验证码获取成功",
			"data":    vc,
		})
	} else if err != nil {
		// 这个是未知错误
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "验证码已发送",
		"data":    nil,
	})
}

// RegisterHandler 用户注册处理器
func RegisterHandler(c *gin.Context) {
	var registerForm form.RegisterDefault
	if err := c.ShouldBind(&registerForm); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "注册信息不能为空",
			"data":    nil,
		})
	}
}

// LoginHandler 登录处理器
func LoginHandler(c *gin.Context) (interface{}, error) {
	var loginForm form.LoginDefault
	if err := c.ShouldBind(&loginForm); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginForm.Username
	password := loginForm.Password
	ok, err := loginPrecheck(username, password)
	if ok {
		return &model.User{
			Username: username,
		}, nil
	}
	return nil, err
}

// loginPrecheck 登录预检查
func loginPrecheck(username string, password string) (bool, error) {
	// 先查询是否存在登录用户
	var loginUser model.User
	config.Db.Where("username = ? or telephone = ? or email = ?", username, username, username).First(&loginUser)
	return true, nil
}
