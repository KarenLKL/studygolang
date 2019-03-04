package persist

import (
	"github.com/KarenLKL/studygolang/crawler/zhenai/parser"
	"log"
)

// item入库处理
func ItemSaver() chan interface{} {
	itemSaverChan := make(chan interface{})
	go func() {
		count := 0
		for {
			item := <-itemSaverChan
			userDetail := parser.ParseUserDetail(item)
			if userDetail != nil {
				log.Printf("save item:got count:#%d,value:%v", count, userDetail)
			}
			count++
		}
	}()
	return itemSaverChan
}
