package main

import (
	"fmt"
	"time"
)

func maker2(c chan int) {
	for e := range c {
		fmt.Println(e)
	}
}

func closeChan(c chan int)  {
	close(c)
}

func main() {
	chanel := make(chan int, 3)
	go maker2(chanel)
	chanel <- 'a'
	chanel <- 'a'
	chanel <- 'a'
	chanel <- 'a'
	closeChan(chanel)
	time.Sleep(time.Millisecond)
}
