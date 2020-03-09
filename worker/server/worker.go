package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/simonzs/crawler_distributed/rpcsupport"
	"github.com/simonzs/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServerRPC(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
