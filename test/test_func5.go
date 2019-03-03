package main

import (
	"fmt"
)

var FunctionForSuffix = map[string]func(string) (string) {
	".gz": GzipFileList, ".tar": TarFileList, ".zip": ZipFileList }

func GzipFileList(ffix string) (string) {
	return ffix
}

func TarFileList(ffix string) (string) {
	return ffix
}

func ZipFileList(ffix string) (string) {
	return ffix
}

func main() {
	function := FunctionForSuffix[".tar"]
	fmt.Print(function(".tar"), "\n")
}