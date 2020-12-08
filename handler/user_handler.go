package handler

import (
	"errors"
	"net/http"

	"github.com/brmyfun/brmy-go/config"
	"github.com/brmyfun/brmy-go/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/brmyfun/brmy-go/model"

	"github.com/brmyfun/brmy-go/form"
	"github.com/brmyfun/brmy-go/util"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// UserInfoHandler 获取用户信息
func UserInfoHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	identityKey := config.Conf.Auth.IdentityKey
	username := claims[identityKey]
	var user model.User
	err := config.Db.Where("username = ?", username).First(&user).Error
	if err == nil {
		c.JSON(http.StatusOK, Ok("查询成功", user))
		return
	}
	c.JSON(http.StatusOK, Err("查询失败"))
	return
}

// RegisterHandler 用户注册
func RegisterHandler(c *gin.Context) {
	var registerForm form.RegisterDefault
	if err := c.ShouldBind(&registerForm); err != nil {
		c.JSON(http.StatusOK, Err("注册信息不能为空"))
		return
	}
	// 判断验证码是否正确
	if vc, err := service.RedisGet(registerForm.Email); err != nil || vc != registerForm.VerificationCode {
		c.JSON(http.StatusOK, Err("验证码错误,请重新发送"))
		return
	}
	var registerUser model.User
	err := config.Db.Where("username = ? or email = ?", registerForm.Username, registerForm.Email).First(&registerUser).Error
	if err == nil {
		c.JSON(http.StatusOK, Err("用户已存在"))
		return
	}
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, Err("查询出错"))
			return
		}
	}
	// 初始化registerUser
	registerUser = model.User{
		Username: registerForm.Username,
		Password: util.Md5Encode(registerForm.Password),
		Email:    registerForm.Email,
	}
	// 插入新用户
	result := config.Db.Create(&registerUser)
	if result.Error != nil {
		c.JSON(http.StatusOK, Err("创建用户出错"))
		return
	}
	c.JSON(http.StatusOK, Ok("用户注册成功", registerUser.ID))
}

// LoginCheck 登录校验
func LoginCheck(username string, password string) (bool, error) {
	// 先查询是否存在登录用户
	var loginUser model.User
	err := config.Db.Where("username = ? or telephone = ? or email = ?", username, username, username).First(&loginUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("用户不存在")
		}
		return false, errors.New("查询出错")
	}
	if !checkPwd(password, loginUser.Password) {
		return false, errors.New("用户名或密码错误")
	}
	return true, nil
}

// checkPwd 密码校验
func checkPwd(password, dbPassword string) bool {
	return util.Md5Encode(password) == dbPassword
}
