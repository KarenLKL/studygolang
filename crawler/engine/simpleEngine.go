package engine

import (
	"fmt"
	"github.com/KarenLKL/studygolang/crawler/fetcher"
	"github.com/KarenLKL/studygolang/crawler/model"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, value := range seeds {
		requests = append(requests, value)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		fmt.Println("fetcher url", r.Url)
		result, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			if value, ok := item.UserInfo.(model.UserInfo); ok {
				fmt.Printf("user info：%v \n", value)
			}
			fmt.Printf("parse city result, the name is：%s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	bytes, err := fetcher.Fetcher(r.Url)
	if err != nil {
		fmt.Printf("fetcher url :%s exception! error:%s \n", r.Url, err.Error())
		return ParseResult{}, err
	}
	return r.ParseFun(bytes, r.Url), nil
}
