package client

import (
	"fmt"

	"github.com/simonzs/crawler_distributed/config"
	"github.com/simonzs/crawler_distributed/rpcsupport"
	"github.com/simonzs/crawler_distributed/worker"
	"github.com/simonzs/crawler_go/engine"
)

// CreateProcessor ...
func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(r engine.Request) (engine.ParserResult, error) {

		sReq := worker.SerializeRequest(r)

		var sResult worker.ParserResult
		err := client.Call(
			config.CrawlServiceRPC, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
