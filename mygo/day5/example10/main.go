package main
//转换json
import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonBuf := `{
	"Company": "zhph",
	"Subjects": [
		"go",
		"go",
		"go"
	],
	"IsOk": true,
	"Price": 666.666
}`
//创建map
m := make(map[string]interface{},4)
err := json.Unmarshal([]byte(jsonBuf),&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	fmt.Printf("%+v\n",m)

	var str string
	for key,value := range m{
		switch data := value.(type) {
		case string:
			str = data
			fmt.Println(str)
		case bool:
			fmt.Println(key,data)
		case float64:
			fmt.Println(key,data)
		case []string:
			fmt.Println(key,data)
		case interface{}:
			fmt.Println("interface:",key,data)
		}
		}

	}



