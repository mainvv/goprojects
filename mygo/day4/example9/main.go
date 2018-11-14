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
	name string
}

func main() {
	var s Student
	s.name = "main" //就近原则。操作Student的那么
	s.Person.name = "go"
	s.sex = 'm'
	s.age = 0
	s.addr = "sh"
	fmt.Printf("s = %+v\n",s)

}
