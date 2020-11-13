package handler

import (
	"errors"
	"net/http"

	"github.com/brmyfun/brmy-go/config"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/brmyfun/brmy-go/model"

	"github.com/brmyfun/brmy-go/form"
	"github.com/brmyfun/brmy-go/util"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// HelloHandler 测试处理器
func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	identityKey := config.Conf.Auth.IdentityKey
	c.JSON(200, gin.H{
		"username": claims[identityKey],
		"text":     "Hello World.",
	})
}

// RegisterHandler 用户注册处理器
func RegisterHandler(c *gin.Context) {
	var registerForm form.RegisterDefault
	if err := c.ShouldBind(&registerForm); err != nil {
		c.JSON(http.StatusOK, Err("注册信息不能为空"))
		return
	}
	// 判断验证码是否正确

	var registerUser model.User
	err := config.Db.Where("username = ? or email = ?", registerForm.Username, registerForm.Email).First(&registerUser).Error
	if err == nil {
		c.JSON(http.StatusOK, Err("用户已存在"))
		return
	}
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, Err("未知错误"))
			return
		}
	}
	// 初始化registerUser
	registerUser = model.User{
		Username: registerForm.Username,
		Password: util.Md5Encode(registerForm.Password),
		Email:    registerForm.Email,
	}
}

// LoginPrecheck 登录预检查
func LoginPrecheck(username string, password string) (bool, error) {
	// 先查询是否存在登录用户
	var loginUser model.User
	err := config.Db.Where("username = ? or telephone = ? or email = ?", username, username, username).First(&loginUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("用户不存在")
		}
		return false, errors.New("查询出错")
	}
	if !CheckPwd(password, loginUser.Password) {
		return false, errors.New("用户名或密码错误")
	}
	return true, nil
}

// CheckPwd 验证密码
func CheckPwd(password, dbPassword string) bool {
	return util.Md5Encode(password) == dbPassword
}
