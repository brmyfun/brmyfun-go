package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/brmyfun/brmy-go/model"
	"github.com/gocolly/colly"
)

// MetaTarget 知乎榜单元素详情
type MetaTarget struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Type          string `json:"type"`
	Created       int64  `json:"created"`
	AnswerCount   string `json:"answer_count"`
	FollowerCount string `json:"follower_count"`
	CommentCount  string `json:"comment_count"`
	Excerpt       string `json:"excerpt"`
}

// MetaChildren 知乎榜单元素缩略图
type MetaChildren struct {
	Thumbnail string `json:"thumbnail"`
}

// ZhihuRankMeta 知乎榜单元素
type ZhihuRankMeta struct {
	DetailText string         `json:"detail_text"`
	Target     MetaTarget     `json:"target"`
	Children   []MetaChildren `json:"children"`
}

// ZhihuRankResp 知乎榜单响应
type ZhihuRankResp struct {
	Data []ZhihuRankMeta `json:"data"`
}

// ZhihuHotRankV1 知乎热榜
func ZhihuHotRankV1() []model.Rank {
	// 入口 https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total
	baseURL := "https://www.zhihu.com/question/"

	var rankResp ZhihuRankResp
	// 初始化 Collector
	c := colly.NewCollector()

	// 定义用于存储榜单的切片
	var rankSlice []model.Rank

	// 获取响应结果
	c.OnResponse(func(resp *colly.Response) {
		// 需要将url换成 https://www.zhihu.com/question/428034871
		json.Unmarshal(resp.Body, &rankResp)
		for key, val := range rankResp.Data {
			rank := model.Rank{
				Type:          "zhihu",
				Title:         val.Target.Title,
				Link:          fmt.Sprint(baseURL, val.Target.ID),
				Excerpt:       val.Target.Excerpt,
				Author:        "知乎热榜",
				Thumbnail:     val.Children[0].Thumbnail,
				Tags:          "",
				Category:      "",
				Metrics:       val.DetailText,
				CommentCount:  val.Target.CommentCount,
				FavoriteCount: "",
				LikeCount:     "",
				AnswerCount:   val.Target.AnswerCount,
				FollowerCount: val.Target.FollowerCount,
				ForwardCount:  "",
				ViewCount:     "",
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

	c.Visit("https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total")

	return rankSlice
}
