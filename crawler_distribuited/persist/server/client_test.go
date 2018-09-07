package main

import (
	"testing"
	"awesomeProject/crawler_distribuited/rpcsupport"
	"fmt"
	"awesomeProject/model"
	"time"
)

func TestRpcServer(t *testing.T) {
	go RpcServer(":9091")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(":9091")
	if err != nil {
		fmt.Println(err)
	}

	profile := model.Profile{
		Age:        23,
		Height:     175,
		Weigth:     120,
		Income:     "2500",
		Gender:     "man",
		Name:       "文龙",
		Xinzuo:     "水平座",
		Occupation: "程序员",
		Marriage:   "未婚",
		House:      "无",
		Hokou:      "广东",
		Education:  "大学本科",
		Car:        "未购车",
	}

	var re string

	err = client.Call("ItemSaverService.Save", profile, &re)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(re)
}
