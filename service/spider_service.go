package service

import (
	"time"

	"github.com/brmyfun/brmy-go/config"
	"github.com/brmyfun/brmy-go/model"
	"github.com/brmyfun/brmy-go/spider"
)

// RunSpider 启动爬虫入口
func RunSpider() {
	// 获取今天的日期
	date := time.Now().Format("2006-01-02")
	// 删除今天的榜单
	config.Db.Where("date = ?", date).Delete(model.Rank{})
}

func zhihuSpider() {
	zhihuRank := spider.ZhihuHotRankV1()
	config.Db.Create(&zhihuRank)
}

func toutiaoSpider() {
	toutiaoRank := spider.ToutiaoHotRankV1()
	config.Db.Create(&toutiaoRank)
}
