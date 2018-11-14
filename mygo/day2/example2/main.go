package main

/**
可变参数
 */
import "fmt"

func main() {
	sum := add(5, 2, 3, 4)
	fmt.Println(sum)
	str := concat("Hello", ",", "World", "!")
	fmt.Println(str)
}

func add(a int, args ... int) int {
	var sum int = a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func concat(a string, str ... string) string {
	var result = a
	//for i := 0;i<len(str);i++{
	//	result += str[i]
	//}
	for _, value := range str {
		result += value
	}
	return result
}
