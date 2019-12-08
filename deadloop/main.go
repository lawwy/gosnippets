package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) //只有一个物理线程的时候，会被死循环占用，无法退出
	go deadloop()
	time.Sleep(time.Second * 5)
	fmt.Println("exit")
}

func deadloop() {
	for {
	}
}
