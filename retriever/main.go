package main

import (
	"fmt"
	"github.com/KarenLKL/studygolang/retriever/mock"
)

type retiever interface {
	Get(str string) string
}

func download(retiever2 retiever) string {
	return retiever2.Get("http://www.imooc.com")
}

func main() {
	var ret retiever
	ret = &mock.Retriever{UserAgent:"chrome"}

	// type assertion
	realret:=ret.(*mock.Retriever)
	fmt.Println(realret.UserAgent)
	//fmt.Println(download(ret))
}
