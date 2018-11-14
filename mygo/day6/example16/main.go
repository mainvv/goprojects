package main
//
import (
	"fmt"
)


func main() {
	ch := make(chan int)//数字通道
	quit := make(chan bool)

	//消费者
	go func() {
		for i := 0; i < 8; i++ {
			num := <- ch
			fmt.Println("num=",num)
		}
		quit <- true
	}()
	//生产者
	Fibonacci(ch,quit)
}

//ch只写 quit只读
func Fibonacci(ch chan<- int, quit <-chan bool) {
	x,y := 1,1
	for{
		//监听channel数据流动
		select {
		case ch <- x:
			x,y = y,x+y
		case flag := <-quit:
			fmt.Println("flag=",flag)
			return
		}
	}
}

