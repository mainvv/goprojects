package main

import "fmt"

var (
	aa = 1
	bb = "csnb"
	//使用并列的方式以类型推断
	cc,dd = 1,true
)
//定义常量
const(
	aaa  = 2333333
	bbb string = "csnb"
)
const(
	ccc = iota
	ddd
	eee
)
const(
	fff = iota
	ggg
	hhh
)
//或者
//const aaa  = 2333333

func callSamePac(){
	fmt.Println("同一包内方法共用")
}

func initValue()  {
	var a int = 10
	var b string = "csnb"
	var c bool = true
	var d,e int = 1,2
	var f,g,h,i = 1,2,"csnb",true
	//var 简写赋值
	j,k,m,n := 1,2,"csnb",true
	fmt.Printf("%d %q %v\n",a, b, c)
	fmt.Printf("%d+%d=%d\n",d, e,d + e)
	fmt.Println(f,g,h,i)
	fmt.Println(j,k,m,n)
	fmt.Println(aaa,bbb)
	fmt.Println(ccc,ddd,eee)
	fmt.Println(fff,ggg,hhh)
}