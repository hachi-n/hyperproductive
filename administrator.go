package hyperproductive

import (
	"sync"
)

type Administrator struct {
	workers []*Worker
	Command
}

func NewAdministrator(numberOfWorker int, command interface{}, params []interface{}) *Administrator {
	administrator := new(Administrator)
	administrator.task = command
	administrator.params = params
	for i := 0; i < numberOfWorker; i++ {
		administrator.AddMember()
	}
	return administrator
}

func (a *Administrator) AddMember() {
	member := NewMember(len(a.workers), a.task, a.params)
	worker := Worker(member)
	a.workers = append(a.workers, &worker)
}

func (a *Administrator) TrustOrder() {
	a.bulkMakeWorkerWork()
}

func (a *Administrator) ExpectOrder() []interface{} {
	return a.bulkIndividualMakeWorkerWork()
}

func (a *Administrator) PrudentOrder() []interface{} {
	return a.bulkPrudentMakeWorkerWork()
}

func (a *Administrator) bulkMakeWorkerWork() {
	wg := new(sync.WaitGroup)
	for _, worker := range a.workers {
		wg.Add(1)
		go makeWorkerWork(worker, wg)
	}
	wg.Wait()
}

func (a *Administrator) bulkIndividualMakeWorkerWork() []interface{} {
	wg := new(sync.WaitGroup)
	requestStream := make(chan interface{}, len(a.workers))
	var result []interface{}

	// producer
	for _, worker := range a.workers {
		wg.Add(1)
		go individualMakeWorkerWork(worker, requestStream)
	}

	// consumer
	go func() {
		for response := range requestStream {
			func() {
				defer wg.Done()
				result = append(result, response)
			}()
		}
	}()
	wg.Wait()
	close(requestStream)

	return result
}

func (a *Administrator) bulkPrudentMakeWorkerWork() []interface{} {
	wg := new(sync.WaitGroup)
	requestStream := make(chan map[int]interface{}, len(a.workers))
	resultMap := make(map[int]interface{})

	// producer
	for _, worker := range a.workers {
		wg.Add(1)
		go prudentMakeWorkerWork(worker, requestStream)
	}

	// consumer
	go func() {
		for response := range requestStream {
			func() {
				defer wg.Done()
				for k, v := range response {
					resultMap[k] = v
				}
			}()
		}
	}()
	wg.Wait()
	close(requestStream)

	// sort
	var result []interface{}
	for i := 0; i < len(a.workers); i++ {
		result = append(result, resultMap[i])
	}

	return result
}
