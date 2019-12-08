package main

import (
	"fmt"
)

/*
	The iota is really an index operator for the current line in the constant declaration block, so if the first use of iota is not the first line in the constant declaration block the initial value will not be zero.
*/

/*
  iota 代表const块中的行的序号（不计空行），从零开始，如果第一行不是iota，则iota则由后续序号开始计算
*/
const (
	azero = iota //0
	aone  = iota //1
)

const (
	info  = "processing"
	bzero = iota //1
	bone  = iota //2
	insfo = "procesxsing"
	bfour = iota //4
)

func main() {
	fmt.Println(azero, aone)
	fmt.Println(bzero, bone, bfour)
}
