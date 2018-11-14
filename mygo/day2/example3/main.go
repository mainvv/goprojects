package main

import "fmt"

func main() {
	var a int = 2
	modify(&a)
	fmt.Println(a)
}

func modify(p *int)  {
	fmt.Println(*p)
	*p = 100
}

func myFunc(a int)  {
	a = 2
	var pa *int= &a
	*pa = 3
	//a值
	fmt.Println(a)
	//a的地址
	fmt.Println(&a)
	//pa值（a的地址）
	fmt.Println(pa)
	//pa的地址
	fmt.Println(&pa)
	//地址的指向值（指针）
	fmt.Println(*pa)
}