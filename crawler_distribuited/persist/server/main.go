package main

import (
	"awesomeProject/crawler_distribuited/rpcsupport"
	"awesomeProject/crawler_distribuited/persist"
	"gopkg.in/olivere/elastic.v2"
	"log"
)

func main() {
	log.Fatal(RpcServer(":9091"))
}

func RpcServer(host string) error {
	//elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://120.78.180.97:9200"),
	)
	if err != nil {
		return err
	}

	//注册服务
	rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client:client,
		})
	return nil
}