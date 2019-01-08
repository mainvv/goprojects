package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("233")

	var _body =233
	fmt.Println(_body)

	/*const (

		a = iota

		b

		c

	)
	fmt.Println(c)*/

	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Print(a & b)
	fmt.Printf("第一行 - c 的值为 %d\n", c )

	c = a | b       /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c )

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:]
	s2 := data[:5]
	copy(s2, s1)
	fmt.Println(data)

	min, max := MinAndMax(33, 22)
	fmt.Printf("min = %d, max = %d\n", min, max)

	x := 10
	y := 20
	swap02(&x, &y)
	fmt.Printf("a = %d, b = %d\n", x, y)
}
//求大小值
func MinAndMax(num1 int, num2 int) (min int, max int) {

	if num1 > num2 {
		min = num2
		max = num1
	} else {
		max = num2
		min = num1
	}
	return
}
func swap01(a, b int) {

	a, b = b, a

	fmt.Printf("swap01 a = %d, b = %d\n", a, b)

}

func swap02(x, y *int) {
	*x, *y = *y, *x
}
