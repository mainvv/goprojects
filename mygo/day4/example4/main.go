package main

import "fmt"

type Student struct{
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	//顺序初始化，每个成员都要初始化
	var s1 Student = Student{1,"main",'m',18,"sh"}
	fmt.Println(s1)
	//指定成员初始化，未初始化的成员复制默认0值
	var s2 Student = Student{name:"张三",addr:"上海"}
	fmt.Println(s2)

	var s3 *Student = &s2
	fmt.Println(*s3)

	p2 := Student{name:"张三",addr:"上海"}
	fmt.Printf("type:%T\n",p2)
	fmt.Printf("type:%T\n",s3)

}
