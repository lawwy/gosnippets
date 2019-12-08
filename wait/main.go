package main

import (
	"fmt"
	"sync"
	"time"
)

//主线程等待子协程完成退出，子协程通过sync.Group告知自身完成状态
func main() {
	group := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		group.Add(1)
		go doWork(i, group)
	}
	group.Wait()
	fmt.Println("done")
}

func doWork(workerId int, group *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	fmt.Println(workerId, "done")
	group.Done() //如果group不是指针，而是值拷贝，则各自操作一个独立的group，引起deadlock
}
