package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(sourceFile, targetFile string) (written int64, err error) {
	srcFile, err := os.Open(sourceFile)
	defer srcFile.Close()
	if err != nil {
		fmt.Printf("open file %s, error: %v", sourceFile, err)
		return 0, err
	}
	reader := bufio.NewReader(srcFile)

	tarFile, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0666)
	defer tarFile.Close()
	if err != nil {
		fmt.Printf("open file %s, error: %v", targetFile, err)
	}
	writer := bufio.NewWriter(tarFile)
	return io.Copy(writer, reader)
}

func main() {
	sourceFile := "/tmp/jiagou.jpg"
	targetFile := "/tmp/jiagou2.jpg"
	written, err := copyFile(sourceFile, targetFile)
	if err != nil {
		fmt.Printf("some error when copy file, error: %v", err)
	} else {
		fmt.Printf("copy file succeed!, copied %v", written)
	}
}
