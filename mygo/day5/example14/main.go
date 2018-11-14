package main
//
import (
	"os"
	"fmt"
	"io"
)

func copyFile(srcFileName,dstFileName string) {

	if srcFileName == dstFileName {
		fmt.Println("源文件和目标文件名字不能相同")
		return
	}
	//只读方式打开源文件
	srcFile,err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	//新建目的文件
	dstFile,err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer srcFile.Close()
	defer dstFile.Close()

	//核心处理，从源文件读取内容往目的文件写
	buf := make([]byte,1024*4)
	for{
		n,err := srcFile.Read(buf)
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			}
		}
		dstFile.Write(buf[:n])
	}
}

func main() {
	copyFile("./day5/example14/a.txt","./test2.txt")
}
