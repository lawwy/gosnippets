package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

func (p data) print2() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	d := data{"one"}
	d.print()
	/*
		只要data这个值可寻址，则可调用指针接受者的方法，实际上会隐式转换。(&d).print()
		但不是所有变量都可寻址，如结构体map的元素不能被寻址（QUESTION）,接口变量也不能被寻址(QUESTION)
	*/
	var p printer = data{"two"} //error
	// var p printer = &data{"two"} //correct
	p.print()

	m := map[string]data{"x": data{"three"}}
	m["x"].print() //error

	m2 := map[string]*data{"x": &data{"three"}}
	m2["x"].print() //correct:结构体指针map中的element能被取址
}
