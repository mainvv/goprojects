package main
//error
import (
	"fmt"
	"errors"
)
//空接口，能保存任意类型的值
type myInter interface {}

type Student struct {
	name string
	id int
}

func myDiv(a,b int)(result int,err error)  {
	if b == 0 {
		return 0,errors.New("分母不能为0")
	}else {
		return a/b,nil
	}
}

func main() {
	err := fmt.Errorf("%s","noraml error")
	fmt.Println(err)

	err2 := errors.New("noraml error2")
	fmt.Println(err2)

	result,err := myDiv(3,0)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
}

