package main

import "fmt"

func main() {
	var data interface{} = "great"
	// 类型转换失败时会返回该类型的零值
	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not a int] value =>", data)
	}
}
