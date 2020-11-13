# 工程依赖

```git
go get -u gorm.io/gorm

go get -u gopkg.in/ini.v1

go get -u github.com/gin-gonic/gin

go get -u github.com/gin-contrib/cors

go get -u github.com/gin-contrib/static

go get -u github.com/appleboy/gin-jwt/v2

go get -u github.com/gocolly/colly/v2

go get -u github.com/axgle/mahonia

go get -u github.com/robfig/cron/v3

go get -u github.com/go-redis/redis/v8

go get -u github.com/matcornic/hermes/v2

go get -u gopkg.in/mail.v2

```

## git 使用方法

```git

git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/brmyfun/brmyfun-go.git
git push -u origin main

```

## spider 爬虫模块

### 模块介绍

使用go语言爬取常用网站的top榜单

### 爬取地址[2020-10-31更新]

1. [知乎热榜API](https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total)
2. [知乎热榜](https://www.zhihu.com/billboard)
3. [知乎热搜榜](https://www.zhihu.com/topsearch)
4. [今日头条热搜榜](https://i.snssdk.com/hot-event/hot-board/?origin=hot_board)
5. [新浪微博热搜榜](https://s.weibo.com/top/summary?cate=realtimehot)
6. [新浪微博话题榜](https://s.weibo.com/top/summary?cate=topicband)
7. [百度热搜榜](http://top.baidu.com/buzz?b=1&fr=topindex)
8. [百度新词榜](http://top.baidu.com/buzz?b=396&fr=topindex)
9. [百度热议榜](http://tieba.baidu.com/hottopic/browse/topicList)
10. [百度手机端热搜榜](http://top.baidu.com/buzz?b=5&fr=topindex)
11. [搜狗微信](https://weixin.sogou.com/)
12. [热点中心](https://article.xmt.cn/api/weibo/platform/hot/event)
13. [阿里热搜榜](https://index.1688.com/alizs/word/listRankType.json?cat=7&rankType=hot&period=week)
