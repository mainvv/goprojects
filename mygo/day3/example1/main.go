package main
/*
冒泡排序
 */
import (
	"math/rand"
	"time"
	"fmt"
)

func main()  {

	rand.Seed(time.Now().UnixNano())
	var a [10]int
	var n = len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	fmt.Println(a)
}
