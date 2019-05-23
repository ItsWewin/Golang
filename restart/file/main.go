package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1Path := "/tmp/abc.txt"
	file2Path := "/tmp/xyz.txt"

	data, err := ioutil.ReadFile(file1Path) // 自动打开或者关闭文件
	if err != nil {
		fmt.Printf("read file error: %v\n", err)
	}

	err = ioutil.WriteFile(file2Path, data, 0666) //自动打开或者关闭文件
	if err != nil {
		fmt.Printf("write file error: %v\n", err)
	}
}
