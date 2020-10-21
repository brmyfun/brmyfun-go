package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(32);not null; comment:'用户名'"`
	Password string `json:"password" gorm:"type:varchar(128);not null; comment:'密码'"`
	Status   int    `json:"status" gorm:"type:TINYINT(1);not null; comment:'状态：1启用 0禁用';default:1"`
}
