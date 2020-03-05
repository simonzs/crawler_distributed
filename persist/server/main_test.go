package main

import (
	"testing"
	"time"

	"github.com/simonzs/crawler_distributed/rpcsupport"
	"github.com/simonzs/crawler_go/engine"
	"github.com/simonzs/crawler_go/model"
)

func TestItemSaver(t *testing.T) {
	// TODO
	// StartItemSaverServer
	go serverRPC(":1234", "test1")
	time.Sleep(time.Second)

	// StartItemSaverClient
	client, err := rpcsupport.NewClient(":1234")
	if err != nil {
		panic(err)
	}
	// Call save

	item := engine.Item{
		URL:  "https://album.zhenai.com/u/1451450381",
		Type: "zhenai",
		ID:   "1451450381",
		Payload: model.Profile{
			Name:       "只等你",
			Gender:     "女士",
			Age:        27,
			Height:     160,
			Income:     "8千-1.2万",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "成都双流区",
		},
	}
	result := ""
	client.Call("ItemSaverService.Save",
		item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s, err: %s", result, err)
	}
}
