package parser

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"regexp"
)

var compile = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]*)</a>`)
var pageLinkCompile = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+/[0-9a-z]+)"[^>]*>([^<]*)</a>`)

func ParseUserList(contents []byte) engine.ParseResult {
	//fmt.Printf("%s",contents)http://album.zhenai.com/u/1618277105
	submatchs := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, submatch := range submatchs {
		item := submatch[2]
		//result.Items = append(result.Items, item)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(submatch[1]),
			ParseFun: func(contents []byte) engine.ParseResult {
				// 利用函数式编程，将用户名称传入到用户信息解析器中
				return ParseUserInfo(contents, string(item))
			},
		})
	}
	pageSubmatch := pageLinkCompile.FindAllSubmatch(contents, -1)
	for _, submatch := range pageSubmatch {
		//result.Items = append(result.Items, item)
		result.Requests = append(result.Requests, engine.Request{
			Url:      string(submatch[1]),
			ParseFun: ParseUserList,
		})
	}
	return result
}
