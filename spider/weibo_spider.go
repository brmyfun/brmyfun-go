package spider

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/brmyfun/brmy-go/model"
	"github.com/gocolly/colly"
)

// WeiboSearchHotRankV1 微博热搜榜
func WeiboSearchHotRankV1() []model.Rank {
	// 入口 https://s.weibo.com/top/summary?cate=realtimehot
	baseURL := "https://s.weibo.com"

	// 初始化 Collector
	c := colly.NewCollector()

	// 定义用于存储榜单的切片
	var rankSlice []model.Rank

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	})

	// 处理HTML
	c.OnHTML(".list_a li", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			key, _ := strconv.Atoi(e.ChildText("strong"))
			link := e.Attr("href")
			metric := e.ChildText("span em")
			titleWithMetric := e.ChildText("span")
			title := regexp.MustCompile(metric).ReplaceAllString(titleWithMetric, "")
			rank := model.Rank{
				Type:          "weibo",
				Title:         title,
				Link:          fmt.Sprint(baseURL, link),
				Excerpt:       "",
				Author:        "微博热搜榜",
				Thumbnail:     "",
				Tags:          "",
				Category:      "",
				Metrics:       metric,
				CommentCount:  "",
				FavoriteCount: "",
				LikeCount:     "",
				AnswerCount:   "",
				FollowerCount: "",
				ForwardCount:  "",
				ViewCount:     "",
				Remark:        "",
				Date:          time.Now().Format("2006-01-02"),
				Rank:          key,
			}
			// 存储数据
			rankSlice = append(rankSlice, rank)
		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("https://s.weibo.com/top/summary?cate=realtimehot")

	return rankSlice
}

// WeiboTopicHotRankV1 微博话题榜
func WeiboTopicHotRankV1() []model.Rank {
	// 入口 https://s.weibo.com/top/summary?cate=topicband
	baseURL := "https://s.weibo.com"

	// 初始化 Collector
	c := colly.NewCollector()

	// 定义用于存储榜单的切片
	var rankSlice []model.Rank

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	})

	// 处理HTML
	c.OnHTML(".list_b li", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			key, _ := strconv.Atoi(e.ChildText("div em"))
			link := e.Attr("href")
			thumbnail := e.ChildAttr("div img", "src")
			title := e.ChildText("article h2")
			excerpt := e.ChildText("article p")
			count := e.ChildText("article span")
			countArr := strings.Fields(count)
			commentCount := regexp.MustCompile("讨论").ReplaceAllString(countArr[0], "")
			viewCount := regexp.MustCompile("阅读").ReplaceAllString(countArr[1], "")
			rank := model.Rank{
				Type:          "weibo",
				Title:         title,
				Link:          fmt.Sprint(baseURL, link),
				Excerpt:       excerpt,
				Author:        "微博话题榜",
				Thumbnail:     thumbnail,
				Tags:          "",
				Category:      "",
				Metrics:       "",
				CommentCount:  commentCount,
				FavoriteCount: "",
				LikeCount:     "",
				AnswerCount:   "",
				FollowerCount: "",
				ForwardCount:  "",
				ViewCount:     viewCount,
				Remark:        "",
				Date:          time.Now().Format("2006-01-02"),
				Rank:          key,
			}
			// 存储数据
			rankSlice = append(rankSlice, rank)
		})
	})

	// 错误处理
	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("https://s.weibo.com/top/summary?cate=topicband")

	return rankSlice
}
