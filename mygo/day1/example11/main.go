package main
/**
Switch示例：猜100以内的随机数
 */
import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	var n int
	//随机数初始化,否则每次
	rand.Seed(time.Now().Unix())
	n = rand.Intn(100)
	//for不加任何参数为死循环
	for{
		var input int
		fmt.Scanf("%d",&input)
		flag := false
		switch {
		case input == n:
			fmt.Println("猜中了")
			flag = true
		case input > n:
			fmt.Println("猜大了")
		case input < n:
			fmt.Println("猜小了")
		}
		if flag{
			break
		}
	}
}
