package main

import (
	"fmt"
	"log"

	"github.com/simonzs/crawler_distributed/config"
	"github.com/simonzs/crawler_distributed/rpcsupport"
	"github.com/simonzs/crawler_distributed/worker"
)

func main() {
	log.Fatal(rpcsupport.ServerRPC(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}))
}