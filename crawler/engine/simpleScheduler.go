package engine

type SimpleScheduler struct {
	workerChan chan Request
}

func (s *SimpleScheduler) WorkerChan() chan Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan Request) {
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan Request)
}

func (s *SimpleScheduler) Submit(r Request) {
	go func() { s.workerChan <- r }()
}
