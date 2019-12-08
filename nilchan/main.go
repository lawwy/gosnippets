package main

import (
	"fmt"
	"time"
)

func main() {
	// nilchandeadlock()
	dynamicInitNilChan()
}

func dynamicInitNilChan() {
	inch := make(chan int)
	outch := make(chan int)
	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int
		for {
			select {
			case val = <-in:
				fmt.Println("get Val from in:", val)
				in = nil
				out = outch
			case out <- val:
				fmt.Println("send Val to out:", val)
				out = nil
				in = inch
			}
		}
	}()
	go func() {
		for r := range outch {
			fmt.Println("result:", r)
		}
	}()
	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}

func nilchandeadlock() {
	var ch chan int
	//未初始化的chan接受跟发送都会阻塞
	//该段程序会引起deadlock
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	fmt.Println("result:", <-ch)
	time.Sleep(2 * time.Second)
}
