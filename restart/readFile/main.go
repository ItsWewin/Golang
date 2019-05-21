package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) // 如果文件不存在则重新创建，使用了打开模式组合 os.O_WRONLY | os.Create
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	// 及时关闭 file 句柄
	defer file.Close()

	str := "hello world!\r\n"

	// 写入时，使用带缓冲区的 *Writer
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为 writer 是带缓冲的，调用 writer.WriteString 实际上是写到缓冲区的，
	//所以需要使用 Flush 方法，将缓冲区的数据写入真正的文件中，否则文件中没有内容
	writer.Flush() // 将内容写入文件
}
