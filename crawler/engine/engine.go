package engine

import (
	"fmt"
	"github.com/KarenLKL/studygolang/crawler/fetcher"
	"github.com/KarenLKL/studygolang/crawler/model"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, value := range seeds {
		requests = append(requests, value)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		fmt.Println("fetcher url", r.Url)
		bytes, err := fetcher.Fetcher(r.Url)
		if err != nil {
			fmt.Printf("fetcher url :%s exception! error:%s \n", r.Url, err.Error())
		}
		result := r.ParseFun(bytes)
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			if value, ok := item.(model.UserInfo); ok {
				fmt.Printf("user info：%v \n", value)
			}
			fmt.Printf("parse city result, the name is：%s", item)
		}
	}
}
