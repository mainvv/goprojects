package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	Person //只有类型，没有变量名--匿名字段来实现继承Person成员
	id int
	addr string
}

func main() {
	//顺序初始化
	var s1 Student = Student{Person{"main",'m',0},0,"sh"}
	fmt.Println(s1)
	fmt.Printf("s1=%+v\n",s1)

	//部分初始化
	s2 := Student{Person:Person{name:"main",age:1},id:0}
	fmt.Printf("s2=%+v\n",s2)

	s2.age =2
	s2.addr ="sh"
	fmt.Printf("s2=%+v\n",s2)

	s2.Person = Person{"go",233,18}
	fmt.Printf("s2=%+v\n",s2)
}
