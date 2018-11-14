package main

import (
	"fmt"
)

func main() {
	var a = 20
	var b = 10
	c := max(a,b)
	fmt.Println(c)
}

func max(num1,num2 int)int{
	if num1 > num2{
		return num1
	}else{
	return num2
	}
}