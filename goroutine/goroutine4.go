package main

import (
	"fmt"
	"sync"
)

type worker struct {
	chanel chan int
	//done   chan bool
	//wg *sync.WaitGroup
	done func()
}

func doWorker(id int, w *worker) {
	for c := range w.chanel {
		fmt.Printf("%c < - chanel id:%d \n", c, id)
		w.done()
	}
}

func createWorker(id int,wg *sync.WaitGroup) *worker {
	w := &worker{
		chanel: make(chan int),
		done:   wg.Done,
	}
	go doWorker(id, w)
	return w
}

func main() {
	var workers [10]*worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i,&wg)
	}

	wg.Add(20)//20个任务
	for i, worker := range workers {
		worker.chanel <- 'a' + i
	}
	for i, worker := range workers {
		worker.chanel <- 'A' + i
	}

	//等待所有任务完成
	wg.Wait()

}
