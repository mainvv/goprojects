package main

import "fmt"
/**
示例：将函数作为参数传值，多返回值函数
 */
type myFunc func(int,int) int

func main(){
	c := sub
	//c := add
	sum := operate(c,4,5)
	fmt.Println(sum)

	//匿名函数
	product := operate(func(a int, b int) int {
		return a * b
	},4,5)
	fmt.Println(product)
	//多返回值函数
	m,n := div(9,4)
	fmt.Println(m,n)
	fmt.Println("------")
	fmt.Println(calculator(4,5,"*"))
	if result,err := calculator(5,4,"/");err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}

func add(a int, b int) int  {
	return a + b
}
func sub(a int, b int)int {
	return a - b
}

func operate(op myFunc, a int, b int) int {
	 return op(a,b)
}

//多返回值函数
func div(a, b int) (int, int) {
	return a/b,a%b
}

func calculator(a,b int,op string) (int,error) {
	switch op {
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		q,_ := div(a, b)
		return q,nil
	default:
		return 0,fmt.Errorf("不支持的运算：" + op)
	}
}