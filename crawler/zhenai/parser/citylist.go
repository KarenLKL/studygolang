package parser

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"regexp"
)

const requestURL = `<a( target="_blank")? href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]*)</a>`

func PrintCityList(contents []byte, _ string) engine.ParseResult {
	compile := regexp.MustCompile(requestURL)
	matches := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for i, item := range matches {
		if i > 2 {
			// 测试时候只取前5个
			break
		}
		//result.Items = append(result.Items, item[3])
		result.Requests = append(result.Requests, engine.Request{Url: string(item[2]), ParseFun: ParseUserList})
	}
	return result
}
