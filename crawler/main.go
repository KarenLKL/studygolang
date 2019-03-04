package main

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/persist"
	"github.com/KarenLKL/studygolang/crawler/scheduler"
	"github.com/KarenLKL/studygolang/crawler/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParseFun: parser.PrintCityList})
	currentEngine := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
		ItemSaver: persist.ItemSaver(),
	}
	currentEngine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParseFun: parser.PrintCityList})

}
