package main

import (
	"fmt"
	"time"
)

func main() {
	var chanels [10]chan<- int

	for i := 0; i < 10; i++ {
		chanels[i] = createMaker(i)
	}
	for i := 0; i < 10; i++ {
		chanels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

//chan<- 规定只能向chan中写入数据
func createMaker(id int) chan<- int {
	chanel := make(chan int)
	go func() {
		fmt.Printf("chanel id:%d value:%c \n", id, <-chanel)
	}()
	return chanel
}
