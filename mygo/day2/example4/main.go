package main

import "fmt"

func main() {
	a := 5
	b := 10
	var x *int = &a
	var y *int = &b
	fmt.Println(*x,*y)
	fmt.Println(&a,&b)
	fmt.Println(*&a,*&b)
	swap(&a,&b)
	fmt.Println(a,b)
}

/*func swap(a, b int) {
	temp := a
	a = b
	b = temp
	return
}*/
func swap(a, b *int) {
	fmt.Println(a,b)
	fmt.Println(*a,*b)
	var temp int= *a
	*a = *b
	*b = temp
	return
}