package main
//
import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close()//关闭标准输出,默认打开

	fmt.Println("main")
	os.Stdout.WriteString("233\n")
	fmt.Println("请输入a")
	var a int
	fmt.Scan(&a)//标准输入到a中
	fmt.Println(a)


}

