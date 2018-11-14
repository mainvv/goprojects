package main

import "fmt"

func main() {
	var a [5]int=[5]int{1,2,3,4,5}
	b := [5]int{6,7,8,9,10}
	c := [5]int{11,12,13}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	c[3] = 14
	c[4] = 15
	fmt.Println(c)

	d := [5]int{1:20,4:30}
	fmt.Println(d)

}
