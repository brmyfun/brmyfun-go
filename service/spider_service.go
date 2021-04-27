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
	// 保存今天的榜单
	zhihuSpider()
	toutiaoSpider()
	weiboSpider()
	baiduSpider()
}

func zhihuSpider() {
	zhihuRank := spider.ZhihuHotRankV1()
	config.Db.Create(&zhihuRank)
}

func toutiaoSpider() {
	toutiaoRank := spider.ToutiaoHotRankV1()
	config.Db.Create(&toutiaoRank)
}

func weiboSpider() {
	weiboSearchRank := spider.WeiboSearchHotRankV1()
	config.Db.Create(&weiboSearchRank)

	weiboTopicRank := spider.WeiboTopicHotRankV1()
	config.Db.Create(&weiboTopicRank)
}

func baiduSpider() {
	baiduSearchRank := spider.BaiduSearchHotRankV1()
	config.Db.Create(&baiduSearchRank)

	baiduTodayRank := spider.BaiduTodayHotRankV1()
	config.Db.Create(&baiduTodayRank)
}
