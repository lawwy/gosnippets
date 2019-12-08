package main

import (
	"fmt"
	"os"
)

func deferCompute() {
	var i int = 1
	//defer 会立即计算参数，相当于defer fmt.Println("result =>",1)
	defer fmt.Println("result =>", func() int { return i }()) //print 1
	i++
}

func deferExcuteAfterFnReturn() {
	files := []string{} //假设有很多
	for _, target := range files {
		f, err := os.Open(target)
		if err != nil {
			fmt.Println("bad target:", target, "error:", err)
		}
		/*
			不能及时关闭，因为defer会在deferEx...函数返回时才执行
			在Unix的文件描述符数量限制下，会报错
			>
				Basically in UNIX platforms, the OS places a limit to the number of open file descriptors that a process may have at any given time.
				The error too many open files is raised since you have reached the limit of file (and or pipe or socket)currently opened and you are trying to open a new file (and or pipe or socket).
				To avoid this problem you must close the file when you have finished to use the open file using the Close() function
			修改方式：
				1. 去掉defer
				2. 新建匿名函数处理打开关闭的逻辑

		*/
		defer f.Close()
	}
}

type Demo struct {
}

func (d *Demo) print(t string) *Demo {
	fmt.Println(t)
	return d
}

func deferWithChain() {
	d := &Demo{}
	defer d.print("1").print("2").print("3")
	d.print("4")
}

func main() {
	// deferCompute()
	deferWithChain()
}
