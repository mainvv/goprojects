package main
//读写文件
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func Writefile(path string) {
	file,err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//使用延迟关闭文件
	defer file.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n",i)
		fmt.Println(buf)
		n,error := file.WriteString(buf)
		if error != nil{
			 fmt.Println(error)
		}
		fmt.Println(n)
	}
}
//读文件
func ReadFile(path string) {
	//打开文件
	file,err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	//延时关闭
	defer file.Close()
	buf := make([]byte,1204*2) //2k大小
	//n表示从文件读取内容的长度
	n,err1 := file.Read(buf)
	if err1 != nil && err1 != io.EOF{ //文件错误，同时没有读到结尾
		fmt.Println("err1=",err1)
		return
	}
	fmt.Println(string(buf[:n]))
}

//每次读取一行
func ReadFileLine(path string)  {
	file,err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//新建缓冲区，把内容放入缓冲区中
	reader := bufio.NewReader(file)
	for {
		buf,err1 := reader.ReadBytes('\n')
		if err1 != nil {
			if err1 == io.EOF {//文件已经结束
				break
			}
			fmt.Println(err1)
		}
		fmt.Println("buf=",string(buf))
	}

}
func main() {
	path := "./test.txt"
	//Writefile(path)
	//ReadFile(path)
	ReadFileLine(path)
}

