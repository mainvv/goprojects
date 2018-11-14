package main

import "fmt"

func main()  {
	sum :=  add(4,5)
	fmt.Println(sum)

	//将函数地址赋值
	c := add
	fmt.Println(c)
	fmt.Println(c(4,5))
}

func add(a int, b int) int  {
	return a + b
}
