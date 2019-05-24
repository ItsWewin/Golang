package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charCount struct {
	chCount    int
	numCount   int
	blankCount int
	otherCount int
}

func main() {
	filePath := "/tmp/test.txt"

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("read file error: %v", err)
		return
	}

	var count charCount
	reader := bufio.NewReader(file)

	for {
		// 读取文件的一行
		// '\n' 是一个字符，不能写成 "\n"
		str, err := reader.ReadString('\n')
		// err 为 io.EOF 说明文件读完
		if err == io.EOF {
			break
		}
		// 遍历行
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'z':
				count.chCount++
			case v == ' ' || v == '\t':
				count.blankCount++
			case v >= 'a' && v <= '9':
				count.numCount++
			default:
				count.otherCount++
			}
		}
	}

	fmt.Printf("char count: %v, number count: %v, blank space count: %v, other count: %v",
		count.chCount, count.numCount, count.blankCount, count.otherCount)
}
