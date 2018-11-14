package list

import "fmt"

var Name string = "2333"
var Age int = 3
func init()  {
	Name = "list2"
	Age = 5
}
func init()  {
	Name = "list"
	Age = 4
}
func List(n int){
	n = Age
	for i:=0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n",i,n-i,n)
	}
}