package main

import (
	"fmt"
	"log"

	"github.com/olivere/elastic"
	"github.com/simonzs/crawler_distributed/config"
	"github.com/simonzs/crawler_distributed/persist"
	"github.com/simonzs/crawler_distributed/rpcsupport"
)

func main() {
	log.Fatal(serverRPC(
		fmt.Sprintf(":%d", config.ItemSaverPort),
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
