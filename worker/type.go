package worker

import (
	"fmt"
	"log"

	"github.com/simonzs/crawler_distributed/config"
	"github.com/simonzs/crawler_go/engine"
	"github.com/simonzs/crawler_go/zhenai/parser"
)

// SerializedParser RPC
type SerializedParser struct {
	Name string
	Args interface{}
}

// Request ...
type Request struct {
	URL    string
	Parser SerializedParser
}

// ParserResult ...
type ParserResult struct {
	Requests []Request
	Items    []engine.Item
}

// SerializeRequest ...
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		URL: r.URL,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

// SerializeResult ...
func SerializeResult(r engine.ParserResult) ParserResult {
	result := ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Reuqests {
		result.Requests = append(
			result.Requests, SerializeRequest(req))
	}
	return result
}

// DeserializeRequest ...
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		URL:    r.URL,
		Parser: parser,
	}, nil

}

// DeserializeResult ...
func DeserializeResult(r ParserResult) engine.ParserResult {
	result := engine.ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing "+
				"request: %v", err)
			continue
		}
		result.Reuqests = append(
			result.Reuqests, engineReq)
	}
	return result
}

// deserializeParser ...
func deserializeParser(
	p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParserCityList:
		return engine.NewFuncParser(
			parser.ParserCityList,
			config.ParserCityList), nil
	case config.ParserCity:
		return engine.NewFuncParser(
			parser.ParserCity,
			config.ParserCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParserProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(
				userName), nil
		} else {
			return nil, fmt.Errorf("invalid "+
				"Arg: %v", p.Args)
		}
	default:
		return nil, fmt.Errorf("unknown parser name: %v", p.Name)
	}
}
