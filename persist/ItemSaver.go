package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v2"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://120.78.180.97:9200"),
	)

	if err != nil {
		panic(err)
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : #%d: %v", itemCount, item)
			itemCount++
			//save to elasticsearch
			_, err := Save(client,item)
			if err != nil {
				log.Printf("Item saver:error"+
					"saving item %v:%v", item, err)
				continue
			}
		}
	}()
	return out
}

func Save(client *elastic.Client, item interface{}) (id string, err error) {
	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").BodyJson(item).Do()

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
