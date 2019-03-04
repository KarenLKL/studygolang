package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	c.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < c.WorkCount; i++ {
		createWorker(in, out)
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

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
