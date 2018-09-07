package main

import (
	"awesomeProject/engine"
	"awesomeProject/zhenai/parser"
	"awesomeProject/scheduler"
	"awesomeProject/crawler_distribuited/persist/client"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:100,
		ItemChan:client.ItemSaver(),
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	/*engine.SimpleEngine{}.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/
}