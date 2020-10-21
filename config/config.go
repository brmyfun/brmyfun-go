package config

import (
	"fmt"
	"log"

	"github.com/brmyfun/brmy-go/model"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conf 全局配置文件
var Conf *Config

// Db 全局数据库连接
var Db *gorm.DB

// Config 配置
type Config struct {
	AppName string `ini:"app_name"`
	AppMode string `ini:"app_mode"`

	Server ServerConfig `ini:"server"`
	MySQL  MySQLConfig  `ini:"mysql"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `ini:"port"`
}

// MySQLConfig MySql配置
type MySQLConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

func init() {
	initConfig()
	initMySQL()
	initTable()
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

func initMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Conf.MySQL.Username, Conf.MySQL.Password, Conf.MySQL.Host, Conf.MySQL.Port, Conf.MySQL.Database)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
}

func initTable() {
	Db.AutoMigrate(&model.User{})
}