package main
/**
for 循环：for的循环条件不能用（）
 */
import "fmt"

func main() {
	printTower(10)
}

func printTower(n int)  {
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}
