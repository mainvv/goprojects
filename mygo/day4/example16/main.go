package main

import "fmt"

type Humaner interface {
	//方法只有声明，没有实现，由别的类型实现
	SayHi()
}

type Student struct {
	name string
	id int
}
type Teacher struct {
	addr string
	group string
}
//Student实现接口方法
func (s *Student)SayHi()  {
	fmt.Println(s)
}
//Teacher实现接口方法
func (t *Teacher)SayHi()  {
	fmt.Println(t)
}

type MyStr string

func (str *MyStr)SayHi()  {
	fmt.Println(*str)
}
//定义一个普通函数，函数的参数为接口类型。只有一个函数，可以有不同表现--多态
func WhoSayHi(i Humaner)  {
	i.SayHi()
}
func main() {
	//var i Humaner
	s := Student{"main",18}
	//i = &s
	WhoSayHi(&s)
	//i.SayHi()
	t := Teacher{"sh","math"}
	//i = &t
	WhoSayHi(&t)
	//i.SayHi()
	var str MyStr = "Hello World"
	//i = &str
	WhoSayHi(&str)
	//i.SayHi()

	x := make([]Humaner,3)
	x[0] = &s
	x[1] = &t
	x[2] = &str

	for _,object := range x{
		object.SayHi()
	}
}

