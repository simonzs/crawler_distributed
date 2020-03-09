package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"

	"github.com/simonzs/crawler_distributed/config"
	itemsaver "github.com/simonzs/crawler_distributed/persist/client"
	"github.com/simonzs/crawler_distributed/rpcsupport"
	worker "github.com/simonzs/crawler_distributed/worker/client"
	"github.com/simonzs/crawler_go/engine"
	"github.com/simonzs/crawler_go/scheduler"
	"github.com/simonzs/crawler_go/zhenai/parser"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_host", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(
		strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)
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

// createClientPool
func createClientPool(
	hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(
			fmt.Sprintf(":%s", h))
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v",
				h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
