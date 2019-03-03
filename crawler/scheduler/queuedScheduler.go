package scheduler

import "github.com/KarenLKL/studygolang/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workChan <- w
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	panic("implement function")
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workChan = make(chan chan engine.Request)
	go func() {
		var (
			requestQ []engine.Request
			workerQ  []chan engine.Request
		)
		for {
			var activeRequestChan engine.Request
			var activeWorkerChan chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequestChan = requestQ[0]
				activeWorkerChan = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workChan:
				workerQ = append(workerQ, w)
			case activeWorkerChan <- activeRequestChan:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
