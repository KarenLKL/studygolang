package engine

type ParseResult struct {
	Items    []Item
	Requests []Request
}

type Item struct {
	Url      string
	Id       string
	Type     string
	UserInfo interface{}
}

type Request struct {
	Url      string
	ParseFun func(contents []byte, url string) ParseResult
}

func NilParseFun(contents []byte) ParseResult {
	return ParseResult{}
}
