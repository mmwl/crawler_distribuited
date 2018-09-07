package persist

import (
	"gopkg.in/olivere/elastic.v2"
	"awesomeProject/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
}

func (i *ItemSaverService) Save(item interface{},result *string) error {
	_, err := persist.Save(i.Client, item)
	if err != nil {
		*result = "fail"
		return err
	}
	*result = "ok"
	log.Printf("Item Saver: %v", item)
	return nil
}
