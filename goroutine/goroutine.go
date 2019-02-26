package main

import (
	"fmt"
	"time"
)

func main() {

	var chanels [10]chan int
	for i := 0; i < 10; i++ {
		chanels[i] = make(chan int)
		go maker(i, chanels[i])
	}

	for i := 0; i < 10; i++ {
		chanels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

func maker(id int, c chan int) {
	fmt.Printf("chanel id:%d value:%c \n", id, <-c)
}
