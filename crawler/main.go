package main

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/front/controller"
	"github.com/KarenLKL/studygolang/crawler/persist"
	"github.com/KarenLKL/studygolang/crawler/scheduler"
	"github.com/KarenLKL/studygolang/crawler/zhenai/parser"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"net/http"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParseFun: parser.PrintCityList})
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	saver := &persist.ItemSaver{
		ElasticClient: client,
		Index:         "dating_profile",
	}
	//go dataParse(saver)
	startWebServer(saver)

}

func dataParse(saver *persist.ItemSaver) {
	items, err := saver.Saver()
	if err != nil {
		panic(err)
	}
	currentEngine := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
		ItemSaver: items,
	}
	currentEngine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParseFun: parser.PrintCityList})
}

func startWebServer(itemSaver *persist.ItemSaver) {
	handler := controller.CreateSearchViewHandler("crawler/front/view/index.html", itemSaver)
	http.Handle("/", http.FileServer(http.Dir("crawler/front/view/")))
	http.Handle("/search", handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Printf("listen http server port 8888 error:%s", err.Error())
		panic(err)
	}
}
