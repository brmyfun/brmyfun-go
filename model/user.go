package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"type:varchar(32);not null;comment:'用户名'"`
	Password  string `json:"password" gorm:"type:varchar(128);not null;comment:'密码'"`
	Telephone string `json:"telephone" gorm:"type:varchar(16);comment:'手机号'"`
	Email     string `json:"email" gorm:"type:varchar(32);comment:'邮箱'"`
	Avatar    string `json:"avatar" gorm:"type:varchar(256);comment:'头像'"`
	Status    int    `json:"status" gorm:"type:TINYINT(1);not null; comment:'状态：1启用 0禁用';default:1"`
}
