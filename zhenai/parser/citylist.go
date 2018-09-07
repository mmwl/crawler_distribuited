/**

	负责解析 text 文本，解析出新的 url 和相关有用的信息

 */

package parser

import (
	"awesomeProject/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`


//将城市列表解析出来
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	results := engine.ParseResult{}

	for _,v := range matches {
		//results.Items = append(results.Items, string(v[2]))
		results.Requests = append(results.Requests, engine.Request{
			Url: string(v[1]),
			ParserFunc: ParseCity,
		})
	}
	return results
}
