package main
//延时功能
import (
	"fmt"
	"time"
)


func main() {
	//创建一个定时器，设置时间为2s，2s后往time通道写入当前时间
/*	timer := time.NewTimer(2*time.Second)
	fmt.Println("当前时间",time.Now())
	//2s后写数据
	t := <- timer.C
	fmt.Println("t=",t)*/

	//只验证1次
/*	timer1 := time.NewTimer(1*time.Second)
	for  {
		<- timer1.C
		fmt.Println("时间到")
	}*/

	//延迟
	<- time.After(2*time.Second)//定时2秒，阻塞2秒，2秒后往channel写内容
	fmt.Println("时间到")
}

