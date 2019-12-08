package main

import (
	"fmt"
	"time"
)

type Result struct {
	Data interface{}
}

type Search func(string) Result

func First(query string, replicas ...Search) Result {
	//无缓冲，说明只有一个能返回，其他会被阻塞，这个场景下会导致资源泄漏
	//资源泄露？是不使用的内存空间泄漏给了其他对象？对象使用了这些本不该被使用的内存空间？
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func First2(query string, replicas ...Search) Result {
	//有足够缓冲，获取第一个，其他会被抛弃
	c := make(chan Result, len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func First3(query string, replicas ...Search) Result {
	c := make(chan Result, 1)
	searchReplica := func(i int) {
		fmt.Println("select ", i+1)
		select {
		//select会随机选择一个可执行的case执行，若没有，则执行default
		//可执行的定义？
		//这里猜测是c是有缓冲的channel，因此发送不会被阻塞（不管执行的方法是否需要等待），因此第一个goroutine会获得发送权，进入第一个case。 而随后的goroutine因为缓冲已满，第一个case将会被阻塞不能执行，因此进入default分支（验证猜想）
		case c <- replicas[i](query):
			fmt.Println("get from ", i+1)
		default:
			//QUESTION:为什么不会走到这里？
			fmt.Println("select default ", i)
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func First4(query string, replicas ...Search) Result {
	c := make(chan Result)
	done := make(chan struct{})
	defer close(done)
	searchReplica := func(i int) {
		select {
		//两个case都会阻塞，此时select会等待其中一个case可执行
		case c <- replicas[i](query):
		case <-done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func MockSearch(id int, delay time.Duration) Search {
	return func(string) Result {
		r := Result{id}
		fmt.Println(id, " sleep ", delay)
		time.Sleep(delay)
		return r
	}
}

func main() {
	searchs := []Search{}
	for i := 1; i < 5; i++ {
		searchs = append(searchs, MockSearch(i, time.Second*time.Duration(i)))
	}
	// r := First3("test", searchs...)
	r := First3("test", searchs...)
	fmt.Println(r)
}
