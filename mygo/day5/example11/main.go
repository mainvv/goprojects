package main
//转换json
import (
	"encoding/json"
	"fmt"
)

type It struct {
	Company string
	Subjects []string
	IsOk bool
	Price float64
}
type It2 struct {
	Subjects []string
}

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
var temp It
err := json.Unmarshal([]byte(jsonBuf),&temp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(temp)
	fmt.Printf("%+v\n",temp)

	var temp2 It2
	err2 := json.Unmarshal([]byte(jsonBuf),&temp2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(temp2)
	fmt.Printf("%+v\n",temp2)
}

