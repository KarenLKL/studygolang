package engine

type SimpleScheduler struct {
	workerChan chan Request
}

func (s *SimpleScheduler) Submit(r Request) {
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(r chan Request) {
	s.workerChan = r
}
