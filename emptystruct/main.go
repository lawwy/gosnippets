package main

import (
	"fmt"
)

type data struct {
}

type data2 struct{}

func main() {
	a := &data{}
	b := &data{}
	c := &data2{}
	if a == b {
		//宽度为零的结构体地址相同，如struct{} ,常用来作为信号传递，如 done := make(chan struct{})
		fmt.Println("same address - a=%p b=%p \n", a, b)
	}
	if c == b {
		//编译出错，不同类型不能==
	}
}
