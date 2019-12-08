package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan struct{})
	jobs := make(chan interface{})
	group := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		group.Add(1)
		//done用于发出停止信号，使得各routine不再进行工作
		//group用于各goroutine汇报工作情况
		go work(i, jobs, done, group)
	}
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(done)
	group.Wait()
	fmt.Println("all done.")
}

func work(workerId int, jobs <-chan interface{}, done <-chan struct{}, group *sync.WaitGroup) {
	defer group.Done()
	for {
		select {
		case job := <-jobs:
			fmt.Printf("%d is doing job:%v ... \n", workerId, job)
		case <-done:
			fmt.Printf("%d receive done signal, do some final things... \n", workerId)
			return
		}
	}
}
