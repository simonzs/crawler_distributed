package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/simonzs/crawler_distributed/config"
	"github.com/simonzs/crawler_distributed/rpcsupport"
	"github.com/simonzs/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRPC(host,
		worker.CrawlService{})
	
	time.Sleep(time.Second)
	// StartCrawlServerClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		URL: "https://album.zhenai.com/u/1451450381",
		Parser: worker.SerializedParser{
			Name: config.ParserProfile,
			Args: "只等你",
		},
	}
	var result worker.ParserResult
	// Call Crawl
	err = client.Call(
		config.CrawlServiceRPC, req, &result)
	if err != nil {
		t.Errorf("err")
	} else {
		fmt.Println(result)
	}
}