package client

import (
	"log"

	"github.com/simonzs/crawler_distributed/rpcsupport"
	"github.com/simonzs/crawler_go/engine"
)

// ItemSaver ...
func ItemSaver(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out

			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err := client.Call("ItemSaverService.Save",
				item, &result)
			if err != nil || result != "ok" {
				log.Print("Item Save: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}
