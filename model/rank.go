package model

import "time"

// RankType 榜单类型
type RankType struct {
	ID        uint      `json:"id" grom:"primaryKey;not null;autoIncrement;comment:'榜单类型ID'"`
	Name      string    `json:"name" gorm:"type:varchar(64);not null;unique;comment:'榜单类型名称'"`
	NameEn    string    `json:"name_en" gorm:"type:varchar(64);not null;unique;comment:'榜单类型英文名称'"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

// Rank 榜单数据
type Rank struct {
	ID            uint      `json:"id" grom:"primaryKey;not null;autoIncrement;comment:'热榜ID'"`
	Type          string    `json:"type" gorm:"type:varchar(64);not null;comment:'热榜类型'"`
	Title         string    `json:"title" gorm:"type:varchar(256);not null;comment:'标题'"`
	Link          string    `json:"link" gorm:"type:varchar(512);not null;comment:'链接'"`
	Excerpt       string    `json:"excerpt" gorm:"type:text;comment:'摘要'"`
	Author        string    `json:"author" gorm:"type:varchar(32);comment:'作者'"`
	Thumbnail     string    `json:"thumbnail" gorm:"type:varchar(512);comment:'缩略图'"`
	Tags          string    `json:"tags" gorm:"type:varchar(128);comment:'标签'"`
	Category      string    `json:"category" gorm:"type:varchar(32);comment:'类别'"`
	Metrics       string    `json:"metrics" gorm:"type:varchar(32);comment:'热度'"`
	CommentCount  string    `json:"commentCount" gorm:"type:varchar(16);comment:'评论数'"`
	FavoriteCount string    `json:"favoriteCount" gorm:"type:varchar(16);comment:'收藏数'"`
	LikeCount     string    `json:"likeCount" gorm:"type:varchar(16);comment:'点赞数'"`
	AnswerCount   string    `json:"answerCount" gorm:"type:varchar(16);comment:'回答数'"`
	FollowerCount string    `json:"followerCount" gorm:"type:varchar(16);comment:'关注数'"`
	ForwardCount  string    `json:"forwardCount" gorm:"type:varchar(16);comment:'转发数'"`
	ViewCount     string    `json:"viewCount" gorm:"type:varchar(16);comment:'浏览数'"`
	Remark        string    `json:"remark" gorm:"type:varchar(256);comment:'备注'"`
	Date          string    `json:"date" gorm:"type:varchar(32);comment:'日期'"`
	Rank          int       `json:"rank" gorm:"type:int;comment:'排名'"`
	CreatedAt     time.Time `json:"createAt"`
}
