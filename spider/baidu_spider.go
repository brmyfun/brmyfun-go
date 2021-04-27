package spider

import (
	"fmt"
	"log"
	"time"

	"github.com/brmyfun/brmy-go/model"
	"github.com/brmyfun/brmy-go/util"
	"github.com/gocolly/colly"
)

// BaiduSearchHotRankV1 百度热搜榜
func BaiduSearchHotRankV1() []model.Rank {
	// 入口 http://top.baidu.com/buzz?b=1&fr=topindex

	// 初始化 Collector
	c := colly.NewCollector()

	// 定义用于存储榜单的切片
	var rankSlice []model.Rank

	// 处理HTML
	c.OnHTML("table.list-table", func(e *colly.HTMLElement) {

		e.ForEach("tr:not(:first-child)", func(key int, e *colly.HTMLElement) {

			// link := e.ChildAttr(".list-title", "href")
			title := e.ChildText(".list-title")
			// metric := e.ChildText(".last")
			// fmt.Printf("title:%s \t link:%s \t metric:%s \n", util.GBK2UTF8(title), link, metric)

			rank := model.Rank{
				Type:          "baidu",
				Title:         util.GBK2UTF8(title),
				Link:          e.ChildAttr(".list-title", "href"),
				Excerpt:       "",
				Author:        "百度热搜榜",
				Thumbnail:     "",
				Tags:          "",
				Category:      "",
				Metrics:       e.ChildText(".last"),
				CommentCount:  "",
				FavoriteCount: "",
				LikeCount:     "",
				AnswerCount:   "",
				FollowerCount: "",
				ForwardCount:  "",
				ViewCount:     "",
				Remark:        "",
				Date:          time.Now().Format("2006-01-02"),
				Rank:          key + 1,
			}
			// 存储数据
			rankSlice = append(rankSlice, rank)

		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("http://top.baidu.com/buzz?b=1&fr=topindex")

	return rankSlice
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

// BaiduTodayHotRankV1 百度今日热点
func BaiduTodayHotRankV1() []model.Rank {
	// 入口 http://top.baidu.com/buzz?b=341&c=513&fr=topbuzz_b1

	// 初始化 Collector
	c := colly.NewCollector()

	// 定义用于存储榜单的切片
	var rankSlice []model.Rank

	// 处理HTML
	c.OnHTML("table.list-table", func(e *colly.HTMLElement) {
		e.ForEach("tr:not(:first-child)", func(key int, e *colly.HTMLElement) {
			// link := e.ChildAttr(".list-title", "href")
			title := e.ChildText(".list-title")
			// metric := e.ChildText(".last")
			// fmt.Printf("title:%s \t link:%s \t metric:%s \n", util.GBK2UTF8(title), link, metric)

			rank := model.Rank{
				Type:          "baidu",
				Title:         util.GBK2UTF8(title),
				Link:          e.ChildAttr(".list-title", "href"),
				Excerpt:       "",
				Author:        "百度热搜榜",
				Thumbnail:     "",
				Tags:          "",
				Category:      "",
				Metrics:       e.ChildText(".last"),
				CommentCount:  "",
				FavoriteCount: "",
				LikeCount:     "",
				AnswerCount:   "",
				FollowerCount: "",
				ForwardCount:  "",
				ViewCount:     "",
				Remark:        "",
				Date:          time.Now().Format("2006-01-02"),
				Rank:          key + 1,
			}
			// 存储数据
			rankSlice = append(rankSlice, rank)
		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("http://top.baidu.com/buzz?b=341&c=513&fr=topbuzz_b1")

	return rankSlice
}
