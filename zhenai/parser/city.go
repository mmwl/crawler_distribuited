package parser

import (
	"awesomeProject/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<th><a href="(.+[^"])" target="_blank">(.+[^<])</a></th>`)

func ParseCity(contents []byte) engine.ParseResult {
	submatch := cityRe.FindAllSubmatch(contents, -1)
	parseResult := engine.ParseResult{}

	for _,v := range submatch{
		request := engine.Request{
			Url: string(v[1]),
			ParserFunc: ParseProfile,
		}
		parseResult.Requests = append(parseResult.Requests,request)
		//名字
		//parseResult.Items = append(parseResult.Items,v[2])
	}

	return parseResult
}

