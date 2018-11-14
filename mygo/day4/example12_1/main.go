package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}
//带有接受者的函数--方法,不支持方法重载
func (person Person) getName() string {
	return person.name
}
//成员赋值
func (person *Person)setName(name string) {
	person.name = name
}

func main() {
	var p Person = Person{"main",'m',18}
	fmt.Println(p.getName())
	p.setName("go")
	//(&p).setName("go")
	fmt.Println(p)
}

