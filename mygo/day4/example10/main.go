package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	//Person //结构体匿名字段
	*Person //指针lei类型结构体匿名字段
	int		//基础类型的匿名字段
	mystr   //自定义类型的匿名字段

}

func main() {
	var s = Student{&Person{"main",'m',0},666,"go"}
	fmt.Println(s.name)
	fmt.Printf("s = %+v\n",s)

	var s2 Student
	s2.Person = new(Person)
	s2.name = "go"
	s2.sex = 'm'
	//s2.int = 233
	//s2.mystr ="233333"
	fmt.Println(s)


}
