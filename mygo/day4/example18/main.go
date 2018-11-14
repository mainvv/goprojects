package main

import "fmt"
//空接口，能保存任意类型的值
type myInter interface {}

func xxx(obj ...myInter)  {
	
}
func main() {
	xxx("aaa")
	xxx(1)
}

