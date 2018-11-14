package main

import "fmt"

func main() {
	//声明map
	var m1 map[int]string
	fmt.Println(m1 == nil)
	//map只有len没有cap.
	fmt.Println(len(m1))
	//通过map创建,可以指定长度，如果没有数据，会显示长度为0
	var m2 = make(map[int]string,2)
	fmt.Println(len(m2))
	m2[0] = "0"
	m2[1] = "1"
	fmt.Println(len(m2))
	//长度超过初始化长度时，会自动增加
	m2[3] = "3"
	fmt.Println(len(m2))

	m3 := map[int]string{0:"0",1:"1",2:"2",3:"3"}
	fmt.Println(m3)

}
