package scheduler

import "github.com/KarenLKL/studygolang/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}
func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}
func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
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
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case activeWorkerChan <- activeRequestChan:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
