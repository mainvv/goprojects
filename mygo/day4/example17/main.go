package main

import "fmt"

type Humaner interface {
	SayHi()
}
//继承
type Personer interface {
	Humaner
	sing(lrc string)
}


type Student struct {
	name string
	id int
}
//Student实现接口方法
func (s *Student)SayHi()  {
	fmt.Printf("Student[%s,%d] say hi\n",s.name,s.id)
}
//Teacher实现接口方法
func (t *Student)sing(lrc string)  {
	fmt.Println("Student在唱歌：",lrc)
}

func main() {
	//
	var i Personer
	s := Student{"main",18}
	i = &s
	i.SayHi()
	i.sing("起来")

	var iPrp Personer
	iPrp = &Student{"go",23}
	var i2 Humaner
	i2 = iPrp //超级可以转换成子集
	i2.SayHi()
	//i2.sing("坐下")
}

