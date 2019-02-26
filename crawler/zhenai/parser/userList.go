package parser

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"regexp"
)

func ParseUserList(contents []byte) engine.ParseResult {
	//fmt.Printf("%s",contents)http://album.zhenai.com/u/1618277105
	compile := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^<]*>([^<]*)</a>`)
	submatchs := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, submatch := range submatchs {
		result.Items = append(result.Items, submatch[2])
		result.Requests = append(result.Requests, engine.Request{Url: string(submatch[1]), ParseFun: ParseUserInfo})
	}
	return result
}
