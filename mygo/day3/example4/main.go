package main

import "fmt"

func main() {
	var a [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}

	//初始化
	b := [3][4]int{
		{1,2,3,4},
		{1,2,3,4},
		{1,2,3,4},
		}
	fmt.Println(b)
}
