package main

import (
	"github.com/simonzs/crawler_distributed/persist/client"
	"github.com/simonzs/crawler_go/engine"
	"github.com/simonzs/crawler_go/scheduler"
	"github.com/simonzs/crawler_go/zhenai/parser"
)

func main() {
	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	// 并发版爬虫架构 Request
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	cityListURL := "https://www.zhenai.com/zhenghun"
	e.Run(engine.Request{
		URL:        cityListURL,
		ParserFunc: parser.ParserCityList,
	})
}
