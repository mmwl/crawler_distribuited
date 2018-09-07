package engine

//需要解析的url以及解析的逻辑
type Request struct {
	Url        string
	Parser     Parser
}

//解析后的结果
type ParseResult struct {
	//又是一个新的需要解析的 Request
	Requests []Request
	//提取出来的东西
	Items    []interface{}
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string,args interface{})
}
type NilParser struct {}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser",nil
}


