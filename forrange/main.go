package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

/*
	方法在调用时等价于
	ff := ({receiverType}).fn
	ff(receiver,arguments...)
	如
	v := field{}
	ff := (*field).print
	print(&v)
*/
func (p *field) print() {
	fmt.Println(p.name)
}

func (p field) print2() {
	fmt.Println(p.name)
}

func rangeStruct() {
	data := []field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		/*
			输出 three,three,three
			for ... range 迭代集合时,将成员值拷贝至v变量中，在整个迭代过程中，v会被复用
			但若v中含有方法（即成员是结构体时），调用方法时会将v进行值拷贝作为方法的接受者调用
			此处等价与
				pp := (*field).print
				go pp(&v)  // 此处传入了v的地址，而v的地址始终不变，因此会输出three,three,three ，注：当使用结构体调用指针接受者的方法时，编译器会自动做这个转换
		*/
		go v.print()
	}
	time.Sleep(time.Second)
	fmt.Println("struct done")
}

func rangePointer() {
	data := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		/*
			输出 one, two ,three (顺序不确定)
			此处等价与
				pp := (*field).print
				go pp(v) // 此处传入了v的值，(该值是指针)，每次迭代v中的值不同，参数进行值拷贝固化下来，因此输出one,two,three
		*/
		go v.print()
	}
	time.Sleep(time.Second)
	fmt.Println("pointer done")
}

func main() {
	rangeStruct()
	rangePointer()
}
