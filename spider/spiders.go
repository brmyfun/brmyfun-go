package spider

import (
	"fmt"
	"log"

	"github.com/brmyfun/brmy-go/util"
	"github.com/gocolly/colly"
)

// WeiboSearchHotRankV1 微博热搜榜
func WeiboSearchHotRankV1() {
	// 入口 https://s.weibo.com/top/summary?cate=realtimehot

	// 初始化 Collector
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	})

	// 处理HTML
	c.OnHTML(".list_a li", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			// url := e.Attr("href")
			// fmt.Println(url)
			title := e.ChildText("span")
			fmt.Println(title)
		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("https://s.weibo.com/top/summary?cate=realtimehot")
}

// WeiboTopicHotRankV1 微博话题榜
func WeiboTopicHotRankV1() {
	// 入口 https://s.weibo.com/top/summary?cate=topicband

	// 初始化 Collector
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	})

	// 处理HTML
	c.OnHTML(".list_b li", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			// url := e.Attr("href")
			// fmt.Println(url)
			imgURL := e.ChildAttr("div img", "src")
			title := e.ChildText("article h2")
			summary := e.ChildText("article p")
			fmt.Println(title, summary, imgURL)
		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("https://s.weibo.com/top/summary?cate=topicband")
}

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
