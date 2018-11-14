package main
//单向管道
import (
	"fmt"
)

//此channel只能写
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}
//此channel只能读
func consumer(ints <-chan int) {
	for num := range ints {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	//生成数据，写入channel
	go producer(ch)
	//消费者，从channe读取内容打印
	consumer(ch)

}

