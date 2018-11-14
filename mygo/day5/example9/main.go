package main
//转换json
import (
	"encoding/json"
	"fmt"
)

type It struct {
	Company string
	//`json:"-"`    此字段不输出
	// `json:"company"`   属性名称小写
	// `json:",string"`   转换成字符串
	Subjects []string
	IsOk bool
	Price float64
}
func main() {
	//将结构体转换成json
	s := It{"zhph",[]string{"go","go","go"},true,666.666}
	//buf,err := json.Marshal(s)
	buf,err := json.MarshalIndent(s,"","	")//格式化转码
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))

	//map转json
	m := make(map[string]interface{},4)
	m["Company"] = "zhph"
	m["Subjects"] = []string{"go","go","go"}
	m["IsOk"] = true
	m["Price"] = 666.666

	result,err := json.MarshalIndent(m,"","	")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}

