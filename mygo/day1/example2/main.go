package main

import (
	"mygo/day1/list"
	"fmt"
)


func main() {
	//同一包类
	callSamePac()
	initValue()
	//不同包类
	list.CallDiffPac()
	list.List(8)
	fmt.Println(list.Name)
	fmt.Println(list.Age)
}
