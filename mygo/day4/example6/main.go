package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}
func main() {
	var s1 Student = Student{0,"main",'m',0,"sh"}
	s2 := Student{0,"main",'m',0,"sh"}
	s3 := Student{0,"main",'m',0,"sh"}
	fmt.Println(s1 == s2)
	fmt.Println(s1 == s3)
	//同类型的两个结构体可以赋值
	var tmp Student = s3
	fmt.Println(tmp)
	//值拷贝
	test1(tmp)
	fmt.Println(tmp)

	test2(&tmp)
	fmt.Println(tmp)
}

/*
地址传递（引用传递）
 */
func test2(student *Student) {
	student.id = 233
	fmt.Println(*student)
}

/*
值传递
 */
func test1(s Student) {
	s.id = 666
	fmt.Println(s)
}