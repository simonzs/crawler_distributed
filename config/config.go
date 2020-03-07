package config

const (

	// Parser Names
	ParserCity     = "ParserCity"
	ParserCityList = "ParserCityList"
	ParserProfile  = "parserProfile"
	NilParser      = "NilParser"

	// Service Ports
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// ElasticSearch
	ElasticIndex = "dating_profile"

	// RPC Endpoints
	ItemSaverRPC    = "ItemSaverService.Save"
	CrawlServiceRPC = "CrawlService.Process"
)
