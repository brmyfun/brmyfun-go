package config

import (
	"log"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

// Conf 全局配置文件
var Conf *Config

// Db 全局数据库连接
var Db *gorm.DB

// Rdb 全局Redis连接
var Rdb *redis.Client

// Config 配置
type Config struct {
	AppName string       `ini:"app_name"`
	AppMode string       `ini:"app_mode"`
	Server  ServerConfig `ini:"server"`
	Auth    AuthConfig   `ini:"auth"`
	MySQL   MySQLConfig  `ini:"mysql"`
	Redis   RedisConfig  `ini:"redis"`
	Email   EmailConfig  `ini:"email"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `ini:"port"`
}

// AuthConfig 鉴权配置
type AuthConfig struct {
	Realm        string `ini:"realm"`
	Key          string `ini:"key"`
	IdentityKey  string `ini:"identity_key"`
	CookieDomain string `ini:"cookie_domain"`
}

// MySQLConfig MySql配置
type MySQLConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// EmailConfig email配置
type EmailConfig struct {
	Server         string `ini:"server"`
	Port           int    `ini:"port"`
	SenderEmail    string `ini:"sender_email"`
	SenderIdentity string `ini:"sender_identity"`
	SMTPUser       string `ini:"smtp_user"`
	SMTPPassword   string `ini:"smtp_password"`
}

func init() {
	initConfig()
}

func initConfig() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Panic("配置文件加载错误!")
	}
	Conf = &Config{}
	err = cfg.MapTo(Conf)
	if err != nil {
		log.Panic("配置文件解析错误!")
	}
}
