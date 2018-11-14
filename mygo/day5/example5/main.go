package main
//字符串转换
import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]byte,0,1024)
	slice = strconv.AppendBool(slice,true)
	//i:追加的数，base：进制
	slice = strconv.AppendInt(slice,233,10)
	slice = strconv.AppendQuote(slice,"main")

	fmt.Println(string(slice))

	var str string
	str = strconv.FormatBool(false)
	fmt.Println(string(str))
	str = strconv.Itoa(666)
	fmt.Println(str)
	//字符串转换成其他类型
	var flag bool
	var err error
	flag,err = strconv.ParseBool("true")
	fmt.Println(flag,err)

	//字符串转换成整形
	a,_ := strconv.Atoi("10")
	fmt.Println(a)



}

