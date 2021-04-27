package main

import (
	"fmt"

	_ "github.com/brmyfun/brmy-go/config"
	"github.com/brmyfun/brmy-go/spider"
)

func main() {
	// router.InitRouter()
	// spider.BaiduTodayHotRankV1()
	fmt.Printf("%v", spider.BaiduTodayHotRankV1())
}
