package worker

import "github.com/simonzs/crawler_go/engine"

// CrawlService ...
type CrawlService struct{}

// Process 接收engine发的消息, 进行反序列化, 在worker, 再序列化
func (CrawlService) Process(
	req Request,
	result *ParserResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
