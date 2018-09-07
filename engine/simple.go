package engine

import (
	"awesomeProject/fetcher"
	"log"
)

type SimpleEngine struct {}

func (s SimpleEngine) Run(seeds ...Request)  {

	var requests []Request
	for _,v := range seeds {
		requests = append(requests, v)
	}

	for len(requests) > 0 {
		//取出队列中的第一个，去网上 Fetch
		r := requests[0]
		requests = requests[1:]

		parserReuest, err := worker(r)
		if err != nil {
			log.Printf("error %s",err)
		}
		//将在网站爬到的 Request 添加进新的队列
		requests = append(requests, parserReuest.Requests...)
		//将爬到的有用信息 item 打印出来
		for _,v := range parserReuest.Items{
			log.Printf("Got item %s",v)
		}
	}
}

func worker(r Request) (ParseResult,error){
	//调用 Fetch 返回原始的网页文本
	body, err := fetcher.Fetch(r.Url)
	//log.Printf("Fetching %s",r.Url)
	if err != nil {
		//如果这条爬虫失败，就爬其他的，服务器不能挂
		log.Printf("Fetcher: error" + "fetching url %s: %v",r.Url,err)
		return ParseResult{},err
	}
	//每个 Request 有特定的解析逻辑, 通过解析网页文本，又解析出很多新的Request
	parserReuest := r.ParserFunc(body)
	return parserReuest,nil
}
