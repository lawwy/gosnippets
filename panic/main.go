package main

import "fmt"

func wrongDemo1() {
	recover() //doesn't do anything
	panic("not good")
	recover() //won't be executed
	fmt.Println("ok")
}

func doRecover() {
	fmt.Println("recovered =>", recover())
}

func wrongDemo2() {
	defer func() {
		doRecover() //panic is not recovered.
	}()
	panic("not good")
}

func goodDemo() {
	defer func() {
		//recover 需要在defer中直接调用才生效，间接调用不生效，参考wrongDemo2
		fmt.Println("recovered:", recover())
	}()
	panic("not good")
}

func main() {
	wrongDemo1()
	wrongDemo2()
	goodDemo()
}
