package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	ReadyNotify
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotify interface {
	WorkerReady(w chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	c.Scheduler.Run()
	for i := 0; i < c.WorkCount; i++ {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, request := range seeds {
		c.Scheduler.Submit(request)
	}
	itemCount := 1
	for {
		result := <-out
		for _, value := range result.Items {
			fmt.Printf("%s, # %d \n", value, itemCount)
			itemCount++
		}
		fmt.Println()

		for _, value := range result.Requests {
			c.Scheduler.Submit(value)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, r ReadyNotify) {
	go func() {
		for {
			r.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
