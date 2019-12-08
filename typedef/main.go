package main

import "sync"

// 为结构体声明别名时不会继承其方法
type myMutex sync.Mutex

// 为接口声明别名时可以保留其定义的方法集
type myLocker sync.Locker

func main() {
	// var mtx myMutex
	// mtx.Lock()
	// mtx.Unlock() //编译出错
	var lock myLocker = new(sync.Mutex)
	lock.Lock()
	lock.Unlock()
}
