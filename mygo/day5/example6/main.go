package main
//正则表达式
import (
	"regexp"
	"fmt"
)

func main() {
	buf := "abc azc a7c aac 888 a9c"
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("error")
		return
	}
	//根据规则提取信息
	result := reg1.FindAllStringSubmatch(buf,-1) //-1表示匹配所有
	fmt.Println(result)

	buf2 := "43.14 567 asdggg 1.23 7. 8.9 111111 6.66 7.88"

	reg2 := regexp.MustCompile(`\d+\.\d+`)
	if reg2 == nil {
		fmt.Println("error")
		return
	}
	result2 := reg2.FindAllStringSubmatch(buf2,-1)
	result3 := reg2.FindAllString(buf2,-1)
	fmt.Println(result2)
	fmt.Println(result3)
}

