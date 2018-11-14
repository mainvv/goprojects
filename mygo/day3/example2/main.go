package main

import "fmt"

func main() {
	//定义一个数组 【3】int 和 【2】int 是不同类型
	var a [3]int
	var b [5]int
	fmt.Println(len(a),len(b))

	//数组长度必须是常量
	//n := 10
	//var c [n]int //non-constant array bound n
	//fmt.Println(len(c))
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	for index,value := range  a{
		fmt.Printf("a[%d]值为:%d\n",index,value)
	}
}
