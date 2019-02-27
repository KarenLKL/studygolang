package engine

type ParseResult struct {
	Items    []interface{}
	Requests []Request
}

type Request struct {
	Url      string
	ParseFun func([]byte) ParseResult
}

func NilParseFun(contents []byte) ParseResult {
	return ParseResult{}
}
