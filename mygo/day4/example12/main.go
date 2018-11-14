package main

import "fmt"

type myint int //
//tmp 接收者,就是传递一个参数
func (tmp myint	 ) Add02(a myint) myint {
	return a + tmp
}

func main() {
	var result int
	result = Add01(1,1)
	fmt.Println(result)

	var r myint = 2
	r = r.Add02(3)
	fmt.Println(r)
}
/**
面向过程
 */
func Add01(a int, b int) int {
	return a + b
}
