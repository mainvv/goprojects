package main
//channel 的close方法
import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("管道增加值：",i)
		}
		close(ch)
		fmt.Println("管道关闭")
		//ch <- 777   //管道关闭后不能再写数据，但是能读数据
	}()

	/*for {
		if num,ok:= <-ch;ok ==true {//如果ok为true，说明管道没有关闭
			fmt.Println("num=",num)
		}else {
			break
		}
	}*/

	for num :=range ch {
		fmt.Println("num=",num)
	}
	
	fmt.Println("读取完毕")

}
