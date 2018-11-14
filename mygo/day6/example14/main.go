package main
//time 定时器
import (
	"fmt"
	"time"
)


func main() {
     timer := time.NewTimer(3*time.Second)
     /*go func() {
     	<- timer.C
     	fmt.Println("son goroutine")
	 }()

	 timer.Stop()//停止定时器
	for {

	}*/

	timer.Reset(time.Second)
	<- timer.C
	fmt.Println("main goroutine")
}

