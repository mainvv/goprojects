package main

import "fmt"

func main() {
	a := [5]int{}
	fmt.Println(len(a))

	s := []int{}
	fmt.Println(s,len(s),cap(s))
	s = append(s,0)
	s = append(s,1)
	s = append(s,2)
	fmt.Println(s,len(s),cap(s))

	//切片的创建
	//1自动推导，初始化
	s1 := []int{1,2,3,4}
	fmt.Println(s1)
	//2 make
	s2 := make([]int,8)
	fmt.Println(s2,len(s2),cap(s2))

	s3 := make([]int,0,1)
	fmt.Println(s3,len(s3),cap(s3))

	//copy
	sa := []int{1,2,3,4}
	sb := []int{6,6,6,6,6,6}
	copy(sa,sb)
	fmt.Println(sa)
	fmt.Println(sb)

}
