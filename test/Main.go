package main

import (
	"awesomeProject/engine"
	"awesomeProject/fetcher"
	"regexp"
	"fmt"
)

//预先编译正则表达式
var nameRe = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">(.+[^<])</h1>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">(.+)</span></td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weigthRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(.+[0-9])KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>(.+[^<])</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>(.+[^<])</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>(.+[^<])</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>(.+[^<])</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>(.+[^<])</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">(.+[^<])</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">(.+[^<])</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">(.+[^<])</span></td>`)
var cityRe = regexp.MustCompile(`<th><a href="(.+[^"])" target="_blank">(.+[^<])</a></th>`)

func main() {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aletai")
	if err != nil {
		return
	}
	city := ParseCity(contents, cityRe)
	for _,v := range city.Items{
		fmt.Printf("%s\n",v)
	}
	for _,v := range city.Requests{
		fmt.Printf("%s\n",v.Url)
	}

}

func ParseCity(contents []byte, re *regexp.Regexp) engine.ParseResult {
	submatch := re.FindAllSubmatch(contents, -1)

	parseResult := engine.ParseResult{}
	for _,v := range submatch{
		request := engine.Request{
			Url: string(v[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
			},
		}
		parseResult.Requests = append(parseResult.Requests,request)
		//名字
		parseResult.Items = append(parseResult.Items,v[2])
	}

	return parseResult
}

