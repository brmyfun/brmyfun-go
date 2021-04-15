package config

import (
	"fmt"
	"log"

	"github.com/brmyfun/brmy-go/model"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	initMySQL()
	initTable()
	initRedis()
	initCron()
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
	Db.AutoMigrate(&model.Rank{})
}

func initRedis() {
	addr := fmt.Sprintf("%s:%s", Conf.Redis.Host, Conf.Redis.Port)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       Conf.Redis.Database,
	})
}

func initCron() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("cron test")
	})
	c.Start()
}
