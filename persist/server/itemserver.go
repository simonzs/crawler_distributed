package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic"
	"github.com/simonzs/crawler_distributed/config"
	"github.com/simonzs/crawler_distributed/persist"
	"github.com/simonzs/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serverRPC(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

// ServerRPC ...
func serverRPC(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServerRPC(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
