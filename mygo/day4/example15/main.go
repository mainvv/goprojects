package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person
	id int
	addr string
}
//继承
func (person Person) PrintInfo() {
	fmt.Println(person)
}
//override 重写
func (s Student)PrintInfo() {
	fmt.Println(s.name,s.sex,s.addr)
}
//成员赋值
func (person *Person)setName(name string) {
	person.name = name
}

func main() {
	//var p Person = Person{"main",'m',18}
	s := Student{Person{"main",'m',18},18,"sh"}
	s.PrintInfo()
	s.Person.PrintInfo()
}

