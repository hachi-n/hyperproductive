package hyperproductive

import (
	"fmt"
	"sync"
)

type Worker interface {
	getHierarchy() int
	Work() interface{}
}

type Command struct {
	task   interface{}
	params []interface{}
}

type Member struct {
	hierarchy int
	Command
}

func NewMember(hierarchy int, command interface{}, params []interface{}) *Member {
	member := new(Member)
	member.hierarchy = hierarchy
	member.task = command
	member.params = params
	return member
}

func (m *Member) getHierarchy() int {
	return m.hierarchy
}

func (m *Member) Work() interface{} {
	var result interface{}
	switch f := m.task.(type) {
	case func():
		f()
	case func(...interface{}):
		f(m.params)
	case func(...interface{}) interface{}:
		result = f(m.params)
	case func() interface{}:
		result = f()
	default:
		// TODO:
		//  Error handling
		fmt.Println("Please Use Type interface{}")
	}
	return result
}

func makeWorkerWork(worker *Worker, wg *sync.WaitGroup) {
	defer wg.Done()
	(*worker).Work()
}

func individualMakeWorkerWork(worker *Worker, requestStream chan interface{}) {
	requestStream <- (*worker).Work()
}

func prudentMakeWorkerWork(worker *Worker, requestStream chan map[int]interface{}) {
	hierarchyResponse := make(map[int]interface{})
	hierarchyResponse[(*worker).getHierarchy()] = (*worker).Work()
	requestStream <- hierarchyResponse
}
