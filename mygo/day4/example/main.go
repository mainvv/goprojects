package main
/*
切片作为参数，传入的是地址的引用，和数组不同
 */
import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//初始化切片
	n := 10
	s := make([]int,n)

	InitDate(s)
	fmt.Println(s)
	sort(s)
	fmt.Println(s)
}

func InitDate(ints []int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(ints); i++ {
		ints[i] = rand.Intn(100)
	}
}
func sort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j],s[j+1] = s[j+1],s[j]
			}
		}
	}
}
