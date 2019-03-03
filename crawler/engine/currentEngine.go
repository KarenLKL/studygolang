package engine

import (
	"fmt"
)

type CurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (c *CurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkCount; i++ {
		createWorker(out, c.Scheduler)
	}

	for _, request := range seeds {
		c.Scheduler.Submit(request)
	}
	itemCount := 1
	for {
		result := <-out
		for _, value := range result.Items {
			fmt.Printf("%s , # %d \n", value, itemCount)
			itemCount++
		}
		fmt.Println()

		for _, value := range result.Requests {
			c.Scheduler.Submit(value)
		}
	}
}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
