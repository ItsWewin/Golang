package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// func ReadFile(filename string) ([]byte, error)
	file := "./test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file error: \n", err)
	}

	fmt.Printf("%v\n", content)         // []byte, 输出: [104 101 108 108 111 32 119 111 114 108 100 10 104 105 10 97 114 101 32 121 111 117 32 111 107 63 10]
	fmt.Printf("%v\n", string(content)) // string, 输出原文内容

	// 没有 Open he close 文件
	// 文件的大家和关闭是被封装在 ioutil.ReadFile 方法中
}
