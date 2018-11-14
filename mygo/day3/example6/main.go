package main

import "fmt"

func main() {
	a := [6]int{1,2,3,4,5,6}
	//数组做参数传递是值传递，其实就是一次值拷贝
	modify(a)
	fmt.Println(a)

	change(&a)
	fmt.Println(a)
}
func change(ints *[6]int) {
	(*ints)[0] = 666
}

func modify(ints [6]int) {
	ints[0] = 666
}

