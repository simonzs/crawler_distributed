package main

import (
	"fmt"

	"github.com/simonzs/crawler_distributed/config"
	itemsaver "github.com/simonzs/crawler_distributed/persist/client"
	worker "github.com/simonzs/crawler_distributed/worker/client"
	"github.com/simonzs/crawler_go/engine"
	"github.com/simonzs/crawler_go/scheduler"
	"github.com/simonzs/crawler_go/zhenai/parser"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	// 并发版爬虫架构 Request
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		ReqeustProcessor: processor,
	}

	cityListURL := "https://www.zhenai.com/zhenghun"
	e.Run(engine.Request{
		URL: cityListURL,
		Parser: engine.NewFuncParser(
			parser.ParserCityList,
			config.ParserCityList),
	})

}
