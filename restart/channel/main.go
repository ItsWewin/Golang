package main

import "fmt"

func main() {
	var intChan chan int
	// 必须要 make，不 make 只声明是没法使用的
	intChan = make(chan int, 2)
	intChan <- 1
	intChan <- 2
	// 容量为 2，最多写入两个元素，超出后会会报错： all goroutines are asleep - deadlock!
	// intChan <- 3
	res := <-intChan
	res2 := <-intChan
	// 最多能读取和写入个数相等的元素，否则会报错: fall goroutines are asleep - deadlock!
	// res3 := <-intChan
	fmt.Printf("res: %v\n", res)
	fmt.Printf("res2: %v\n", res2)
	// fmt.Printf("res3: %v\n", res3)
}
