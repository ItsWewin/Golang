package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查询当前服务器逻辑 cpu 数量
	cpuNum := runtime.NumCPU()
	fmt.Printf("cpu num: %v\n", cpuNum)

	// 设置程序运行时使用的 cpu 数量
	runtime.GOMAXPROCS(cpuNum - 1)
	// go1.8 之前，需要设置下程序运行时使用的 cpu 的数量，可以提高效率
}
