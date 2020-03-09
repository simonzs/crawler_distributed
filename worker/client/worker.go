package client

import (
	"net/rpc"

	"github.com/simonzs/crawler_distributed/config"
	"github.com/simonzs/crawler_distributed/worker"
	"github.com/simonzs/crawler_go/engine"
)

// CreateProcessor ...
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(r engine.Request) (engine.ParserResult, error) {

		sReq := worker.SerializeRequest(r)

		var sResult worker.ParserResult
		c := <-clientChan
		err := c.Call(
			config.CrawlServiceRPC, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
