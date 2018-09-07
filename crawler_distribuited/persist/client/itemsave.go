package client

import (
	"log"
	"awesomeProject/crawler_distribuited/rpcsupport"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	client, err := rpcsupport.NewClient(":9091")
	if err != nil {
		panic(err)
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : #%d: %v", itemCount, item)
			itemCount++
			var result string
			//save to elasticsearch
			err := client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("Item saver:error"+
					"saving item %v:%v", item, err)
				continue
			}
		}
	}()
	return out
}
