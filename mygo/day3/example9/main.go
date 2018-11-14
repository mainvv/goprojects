package main

import "fmt"

func main() {
	array := []int{0,1,2,3,4,5,6,7,8,9}
	//切片截取
	//[low:high:max]
	s1 := array[:]
	fmt.Println(s1,len(s1),cap(s1))

	data := array[2]
	fmt.Println(data)

	s2 := array[3:6:7]
	fmt.Println(s2,len(s2),cap(s2))

	s3 := array[:6]
	fmt.Println(s3,len(s3),cap(s3))

	s4 := array[3:]
	fmt.Println(s4,len(s4),cap(s4))

	a := []int{0,1,2,3,4,5,6,7,8,9}
	x1 := a[2:5]
	x1[1] = 666
	fmt.Println(x1)

	x2 := x1[2:7]
	x2[2] = 777
	fmt.Println(x2)

}
