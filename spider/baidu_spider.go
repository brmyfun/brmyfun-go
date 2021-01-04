package spider

import (
	"fmt"
	"log"

	"github.com/brmyfun/brmy-go/util"
	"github.com/gocolly/colly"
)

// BaiduSearchHotRankV1 百度热搜榜
func BaiduSearchHotRankV1() {
	// 入口 http://top.baidu.com/buzz?b=1&fr=topindex

	// 初始化 Collector
	c := colly.NewCollector()

	// 处理HTML
	c.OnHTML("table.list-table", func(e *colly.HTMLElement) {
		e.ForEach(".list-title", func(_ int, e *colly.HTMLElement) {
			title := e.Text
			fmt.Printf("title:%s \n", util.GBK2UTF8(title))
		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("http://top.baidu.com/buzz?b=1&fr=topindex")
}

// BaiduNewWordHotRankV1 百度新词榜
func BaiduNewWordHotRankV1() {
	// 入口 http://top.baidu.com/buzz?b=396&fr=topindex

	// 初始化 Collector
	c := colly.NewCollector()

	// 处理HTML
	c.OnHTML("table.list-table", func(e *colly.HTMLElement) {
		e.ForEach(".list-title", func(_ int, e *colly.HTMLElement) {
			title := e.Text
			fmt.Println(util.GBK2UTF8(title))
		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("http://top.baidu.com/buzz?b=396&fr=topindex")
}
