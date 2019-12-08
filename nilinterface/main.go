package main

import "fmt"

func nilInterface() {
	var data *byte
	var in interface{} //接口值由类型跟值组成，只有类型跟值均为空，in == nil 才成立

	fmt.Println(data, data == nil) //<nil> true :
	fmt.Println(in, in == nil)     //<nil> true : 接口的type跟value都为空

	in = data
	fmt.Println(in, in == nil) //<nil> false : in接口type不为空，value为空
}

func wrongNilReturn() {
	doit := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		}
		return result
		// return nil //correct:显式地返回nil
	}
	if res := doit(-1); res != nil {
		//这里res的type不为空，value为空
		fmt.Println("good result:", res)
	}
}

func main() {
	nilInterface()
	wrongNilReturn()
}
