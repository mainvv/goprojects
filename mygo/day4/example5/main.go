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
	var s Student = Student{id:0,name:"main"}
	fmt.Println(s)
	//操作结构体成员用.
	s.id = 1
	s.name = "11"
	s.sex = 'm'
	s.age = 18
	s.addr = "sh"
	fmt.Println(s)

	//指针有合法指向后才能操作成员,p1.id和(*p1).id等价
	var s1 Student
	var p1 *Student= &s1
	fmt.Println(p1)
	p1.id = 2
	p1.sex ='2'
	(*p1).name = "22"
	fmt.Println(p1)

	//2.通过new申请结构体
	p2 := new(Student)
	p2.id = 3
	p2.name ="33"
	p2.sex = 'm'
	fmt.Println(p2)

}
