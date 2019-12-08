package main

import (
	"fmt"
	"reflect"
)

//如果结构体中的元素都能用==比较，则该结构体可以用==比较，否则会产生编译错误
type comparable struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

type unComparable struct {
	num    int
	checks [10]func() bool   // not comparable
	doit   func() bool       // not comparable
	m      map[string]string // not comparable!!注意
	bytes  []byte            // not comparable!!注意
}

func main() {
	unCompare()
	compare()
}

func unCompare() {
	v1 := unComparable{}
	v2 := unComparable{}
	// fmt.Println("v1 == v2:", v1 == v2) //编译错误
	fmt.Println("v1==v2", reflect.DeepEqual(v1, v2)) //对于无法用==比较的struct可使用DeepEqual比较
}

func compare() {
	v1 := comparable{}
	v2 := comparable{}
	fmt.Println("v1==v2:", v1 == v2) //true

	v3 := comparable{
		events: make(<-chan string),
	}
	v4 := comparable{
		events: make(<-chan string),
	}
	fmt.Println("v3==v4:", v3 == v4)
	//false
}
