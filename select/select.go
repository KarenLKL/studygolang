package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	chanel := make(chan int)
	go func() {
		var i = 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			chanel <- i
			i++
		}
	}()
	return chanel
}

func walk(id int, c chan int) {
	for e := range c {
		fmt.Printf("walk id:%d, value: %d \n", id, e)
	}
}

func createWorker(id int) chan int {
	chanel := make(chan int)
	go walk(id, chanel)
	return chanel
}

func main() {
	c1, c2 := generator(), generator()
	var worker = createWorker(0)
	var activeWorker chan int //nil chan

	var values []int
	activeValue, ticker, after := 0, time.NewTicker(time.Second), time.After(10*time.Second)
	for {
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case v := <-c1:
			values = append(values, v)
		case v := <-c2:
			values = append(values, v)
		case activeWorker <- activeValue:
			if len(values) > 0 {
				values = values[1:]
			}
		case <-ticker.C:
			fmt.Printf("len（values）= %d", len(values))
		case <-time.After(800 * time.Millisecond):
			// 超过800毫秒，超时
			fmt.Println("timeout...")
		case <-after:
			// 10秒钟后退出程序
			fmt.Println("Bey Bey!")
			return

		}

	}
}
