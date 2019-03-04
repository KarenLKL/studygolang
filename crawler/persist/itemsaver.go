package persist

import (
	"log"
)

// item入库处理
func ItemSaver() chan interface{} {
	itemSaverChan := make(chan interface{})
	go func() {
		count := 0
		for {
			item := <-itemSaverChan
			log.Printf("save item:got count:#%d,value:%v", count, item)
			count++
		}
	}()
	return itemSaverChan
}
