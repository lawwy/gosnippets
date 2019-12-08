package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cancel := make(chan struct{})
	group := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		group.Add(1)
		worker := NewWorker(i, cancel, group)
		go worker.Do()
		// //or:
		// go work(i, cancel, group)
	}
	close(cancel)
	group.Wait()
	fmt.Println("done")
}

func work(workId int, cancel <-chan struct{}, group *sync.WaitGroup) {
	fmt.Println(workId, " running...")
	time.Sleep(time.Second * 3)
	defer group.Done()
	<-cancel
	fmt.Println(workId, " receive cancel, do some final things")
}

//封装成worker
type Worker struct {
	id     int
	cancel <-chan struct{}
	group  *sync.WaitGroup
}

func (worker *Worker) Do() {
	fmt.Println(worker.id, " running...")
	time.Sleep(time.Second * 3)
	defer worker.group.Done()
	<-worker.cancel
	fmt.Println(worker.id, " receive cancel, do some final things")
}

func NewWorker(workId int, cancel <-chan struct{}, group *sync.WaitGroup) *Worker {
	return &Worker{workId, cancel, group}
}
