package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := [5]int{1,2,3,4,5}
	b := [5]int{1,2,3,4,5}
	c := [5]int{1,2,3,4,6}
	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(a != c)

	//设置种子
	//rand.Seed(666)
	rand.Seed(time.Now().UnixNano())//以当前系统时间作为种子
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int(),rand.Intn(100))
	}
}

