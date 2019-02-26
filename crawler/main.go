package main

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParseFun: parser.PrintCityList})
}
