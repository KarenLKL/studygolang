package main

import (
	"html/template"
	"net/http"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParseFun: parser.PrintCityList})
	//items, err := persist.ItemSaver("dating_profile")
	//if err != nil {
	//	panic(err)
	//}
	//currentEngine := &engine.ConcurrentEngine{
	//	Scheduler: &scheduler.QueuedScheduler{},
	//	WorkCount: 100,
	//	ItemSaver: items,
	//}
	//currentEngine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParseFun: parser.PrintCityList})
	startWebServer()

}

func startWebServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("crawler/front/index.html")
		if err != nil {
			panic(err)
		}
		_ = t.Execute(writer, "hello world!")
	})
	_ = http.ListenAndServe(":8888", nil)
}
