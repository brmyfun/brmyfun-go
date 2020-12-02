package model

import (
	"time"
)

// User 用户
type User struct {
	ID        uint      `json:"id" grom:"primaryKey;not null;autoIncrement;comment:'用户ID'"`
	Username  string    `json:"username" gorm:"type:varchar(32);not null;unique;comment:'用户名'"`
	Password  string    `json:"-" gorm:"type:varchar(128);not null;comment:'密码'"`
	Telephone string    `json:"telephone" gorm:"type:varchar(16);comment:'手机号'"`
	Email     string    `json:"email" gorm:"type:varchar(32);comment:'邮箱'"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(256);comment:'头像'"`
	Status    int       `json:"status" gorm:"type:tinyint(1);not null; comment:'状态：1启用 0禁用';default:1"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

// APIUser 接口返回用户
type APIUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
