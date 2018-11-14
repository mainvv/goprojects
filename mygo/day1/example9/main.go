package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	const fileName = "day1/example9/a.txt"
	content,error := ioutil.ReadFile(fileName)
	if error != nil{
		fmt.Println(error)
	}else{
		fmt.Printf("%s\n",content)
	}
}
