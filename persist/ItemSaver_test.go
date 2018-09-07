package persist

import (
	"testing"
	"awesomeProject/model"
	"gopkg.in/olivere/elastic.v2"
	"encoding/json"
	"fmt"
)

func TestSave(t *testing.T) {

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

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://120.78.180.97:9200"),
	)

	id, err := Save(client,profile)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://120.78.180.97:9200"),
	)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).Do();
	if err != nil {
		panic(err)
	}

	var actual model.Profile

	err = json.Unmarshal([]byte(*resp.Source),&actual)

	if err != nil {
		panic(err)
	}

	fmt.Println("%v",actual)
}
