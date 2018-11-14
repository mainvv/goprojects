package main

import "fmt"

/*
slice切片
 */
func main() {
	a := []int{1,2,3,4,5}
	/*	切片初始化
		2 : 起点
		3 ： 终点
	前闭后开，长度：终点 - 起点
		4 ： max
	容量：max - 起点*/
	s := a[1:3:4]
	fmt.Println(s)
	fmt.Println(len(s)) //长度
	fmt.Println(cap(s)) //容量
}

