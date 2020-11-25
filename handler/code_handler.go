package handler

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/brmyfun/brmy-go/config"

	"github.com/brmyfun/brmy-go/service"
	"github.com/brmyfun/brmy-go/util"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

// TelephoneVcHandler 发送手机验证码处理器
func TelephoneVcHandler(c *gin.Context) {
	telephone := c.PostForm("telephone")
	if !util.VerifyTelephone(telephone) {
		c.JSON(http.StatusOK, Err("手机号格式有误"))
		return
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
		// 发送手机验证码逻辑 todo
		c.JSON(http.StatusOK, Ok("验证码发送成功", vc))
		return
	} else if err != nil {
		// 这个是未知错误
		panic(err)
	}
	c.JSON(http.StatusOK, Ok("验证码已发送", vc))
}

// EmailVcHandler 发送邮箱验证码处理器
func EmailVcHandler(c *gin.Context) {
	email := c.PostForm("email")
	if !util.VerifyEmail(email) {
		c.JSON(http.StatusOK, Err("邮箱地址格式有误"))
		return
	}
	// 先从缓存中获取验证码
	vc, err := service.RedisGet(email)
	if err == redis.Nil {
		// 缓存中没有则生成一个新的验证码
		vc = util.GenRandomStrCode(6)
		// 存入缓存
		_, err := service.RedisSet(email, vc, time.Minute)
		if err != nil {
			// 这个错误需要处理一下
			c.JSON(http.StatusOK, Err("内部错误"))
			return
		}
		// 发送邮箱验证码逻辑
		if err := sendEmail(email, vc.(string)); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, Err("验证码发送失败"))
			return
		}
		c.JSON(http.StatusOK, Ok("验证码发送成功", vc))
		return
	} else if err != nil {
		// 这个是未知错误
		c.JSON(http.StatusOK, Err("内部错误"))
		return
	}
	c.JSON(http.StatusOK, Ok("验证码已发送", vc))
}

// sendEmail 发送邮件
func sendEmail(email, code string) error {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      "白日梦语",
			Link:      "https://brmy.fun",
			Logo:      "https://ftp.bmp.ovh/imgs/2020/11/fffb97d83cf6819b.png",
			Copyright: "Copyright © 2020 brmyfun. All rights reserved.",
		},
	}
	template := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"终于等到你啦！感谢你注册白日梦语账号。",
			},
			Actions: []hermes.Action{
				{
					Instructions: "这是你的注册验证码：",
					InviteCode:   code,
				},
			},
			Signature: "Thanks",
		},
	}
	emailBody, err := h.GenerateHTML(template)
	if err != nil {
		panic(err)
	}

	from := mail.Address{
		Name:    config.Conf.Email.SenderIdentity,
		Address: config.Conf.Email.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", email)
	m.SetHeader("Subject", "注册验证码")
	m.SetBody("text/html", emailBody)
	fmt.Println(config.Conf.Email.Server, config.Conf.Email.Port, config.Conf.Email.SMTPUser, config.Conf.Email.SMTPPassword)
	d := gomail.NewDialer(config.Conf.Email.Server, config.Conf.Email.Port, config.Conf.Email.SMTPUser, config.Conf.Email.SMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(m)
}
