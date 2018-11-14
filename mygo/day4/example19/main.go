package main
//本示例实现断言
import (
	"fmt"
)
//空接口，能保存任意类型的值
type myInter interface {}

type Student struct {
	name string
	id int
}

func myFunc(obj ...myInter)  {
	
}
func main() {
	i := make([]interface{},3)
	i[0] = 1
	i[1] = "main"
	i[2]= Student{"main",233}
	for _,data := range i{
		//data.(类型) 来实现断言
		if value,ok := data.(int);ok == true{
			fmt.Println(ok,value)
		}else if value, ok := data.(string);ok == true {
			fmt.Println(ok,value)
		}else if value,ok := data.(Student);ok ==true{
			fmt.Println(ok,value)
		}
		//myFunc(data)
	}

	//switch
	for index,data := range i{
		switch value := data.(type){
		case int:
			fmt.Println(index,value)
		case string:
			fmt.Println(index,value)
		case Student:
			fmt.Println(index,value)

		}
	}
}

