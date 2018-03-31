package service

import (
	"sync"
)

type WorkerService struct {
	workerNum int
	waitGroup *sync.WaitGroup
	channel   chan interface{}
}

func CreateWorkerService(workerNum int, executor func(interface{})) *WorkerService {
	service := &WorkerService{workerNum: workerNum}
	service.waitGroup = &sync.WaitGroup{}
	service.channel = make(chan interface{}, workerNum*5)
	for i := 0; i < workerNum; i++ {
		go service.worker(service.channel, executor)
	}
	return service
}

func (self *WorkerService) Submit(value interface{}) {
	self.waitGroup.Add(1)
	self.channel <- value
}

func (self *WorkerService) worker(values <-chan interface{}, executor func(interface{})) {
	for value := range values {
		executor(value)
		self.waitGroup.Done()
	}
}

func (self *WorkerService) Wait() {
	if self.waitGroup != nil {
		self.waitGroup.Wait()
	}
}
