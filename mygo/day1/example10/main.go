package main

import "fmt"

func main()  {
	//test1()
	g := grad(100)
	fmt.Println(g)
}

func test1() {
	var a int = 0
	//区别于java的break，不加自动中断
	//用fallthrough继续执行下一个case
	switch a {
	case 0,1,2,3,4:
		fmt.Println(0)
		fallthrough
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("default")
	}
}

func grad(score int)string  {
	g := ""
	switch {
	case score < 0 || score >100:
		panic(fmt.Sprintf("wrong score:%d",score))
	case score < 60:
		g = "c"
	case score < 80:
		g = "b"
	case score <= 100:
		g = "a"
	}
	return g
}