package persist

import (
	"log"

	"github.com/olivere/elastic"
	"github.com/simonzs/crawler_go/engine"
	"github.com/simonzs/crawler_go/persist"
)

// ItemSaverService ...
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// Save ...
func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v",
			item, err)
	}
	return err
}
