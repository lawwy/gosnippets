package main

import "fmt"

type data struct {
	name string
}

func errInvalidMemAddress() {
	m := map[string]*data{"x": {"one"}}
	m["z"].name = "what?" //对象未初始化，非法指针
	//panic: runtime error: invalid memory address or nil pointer dereference
}

func successAssignStructPtrMapElement() {
	m := map[string]*data{
		"x": {"one"}, //用简略形式初始化时不用取址符&？
	}
	m["x"].name = "two" //结构体指针map的元素能取址，of cource
	fmt.Println(m["x"])
}

func errAssignStructMapElement() {
	m := map[string]data{"x": {"one"}}
	// m["x"].name = "two" //结构体map的元素不能取址，报错
	// 似乎是因为低层map可能会发生扩容的情况，迁移键值对，原来的地址会失效，考虑到安全，不允许这样操作

	//需要用一个局部变量才能赋值
	r := m["x"]
	r.name = "two"
	fmt.Println("%v", m) // map[x:{one}], ?QUESTION: m的元素不能被取址，因此r是值的拷贝？
	m["x"] = r
	fmt.Println("%v", m) // map[x:{two}]
}

func successAssignSliceElement() {
	s := []data{{"one"}}
	s[0].name = "two" //correct，结构体slice的元素能被取址
	fmt.Println(s)
}

func main() {
	// errAssignStructMapElement()
	// successAssignSliceElement()
	// test()
	errInvalidMemAddress()
}
