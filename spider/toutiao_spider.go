package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/brmyfun/brmy-go/model"
	"github.com/gocolly/colly"
)

// ToutiaoRankMeta 头条榜单元素
type ToutiaoRankMeta struct {
	Title    string `json:"Title"`
	HotValue string `json:"HotValue"`
}

// ToutiaoRankResp 头条榜单响应
type ToutiaoRankResp struct {
	Data []ToutiaoRankMeta `json:"data"`
}

// ToutiaoHotRankV1 今日头条热榜
func ToutiaoHotRankV1() []model.Rank {
	// 入口 https://i.snssdk.com/hot-event/hot-board/?origin=hot_board
	baseURL := "https://www.toutiao.com/search/?keyword="

	var rankResp ToutiaoRankResp

	// 初始化 Collector
	c := colly.NewCollector()

	// 定义用于存储榜单的切片
	var rankSlice []model.Rank

	// 获取响应结果
	c.OnResponse(func(resp *colly.Response) {
		// 需要将link换成 https://www.toutiao.com/search/?keyword=title
		json.Unmarshal(resp.Body, &rankResp)

		for key, val := range rankResp.Data {
			rank := model.Rank{
				Type:          "toutiao",
				Title:         val.Title,
				Link:          fmt.Sprint(baseURL, val.Title),
				Excerpt:       "",
				Author:        "头条热搜",
				Thumbnail:     "",
				Tags:          "",
				Category:      "",
				Metrics:       val.HotValue,
				CommentCount:  0,
				FavoriteCount: 0,
				LikeCount:     0,
				AnswerCount:   0,
				FollowerCount: 0,
				ForwardCount:  0,
				ViewCount:     0,
				Remark:        "",
				Date:          time.Now().Format("2006-01-02"),
				Rank:          key + 1,
			}
			// 存储数据
			rankSlice = append(rankSlice, rank)
		}
	})

	c.OnError(func(resp *colly.Response, err error) {
		log.Println("Request URL:", resp.Request.URL, "failed with response:", resp, "\nError:", err)
	})

	c.Visit("https://i.snssdk.com/hot-event/hot-board/?origin=hot_board")

	return rankSlice
}
