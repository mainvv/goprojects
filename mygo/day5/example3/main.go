package main
//panic  defer  recover
import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaa")
}
func testb(x int) {
	defer func() {
		if err := recover();err != nil{
			fmt.Println(err)
		}
	}()
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
	testb(2)
	testc()

}

