package main
//panic
import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaa")
}
func testb(x int) {
	fmt.Println("bbbbbbb")
	//显示调用panic，导致程序中断
	//panic("panic")
	var a [10]int
	a[x] = 233

}
func testc() {
	fmt.Println("ccccccc")
}

func main() {
	testa()
	testb(11)
	testc()

}

