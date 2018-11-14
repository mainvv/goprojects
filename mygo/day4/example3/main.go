package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{0:"0",1:"1",2:"2",3:"3"}
	fmt.Println(m1)
	m1[1] = "11"//修改
	fmt.Println(m1)
	m1[4] = "4"//追加，map底层自动扩容，和append类似
	fmt.Println(m1)

	//遍历map，第一个返回值为key，第二个返回值value
	for key,value := range m1{
		fmt.Println(key,value)
	}

	//判断key是否存在
	value,IsExist := m1[4]

	if IsExist {
		fmt.Println(value)
	}else {
		fmt.Println("不存在")
	}

	//删除map里的键值对
	delete(m1,1)
	fmt.Println(m1)

	//map作为参数为引用传递
	del(m1)
	fmt.Println(m1)
}
func del(strings map[int]string) {
	delete(strings,2)
}
