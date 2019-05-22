package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发计算 1 - 200 各个数据的阶乘，并把结果存放到 map 中
var (
	factResult = make(map[int]int)

	//声明一个全局互斥锁
	//lock 是一个全局互斥锁
	//sync 是内置包
	// Mutex: 互斥
	lock sync.Mutex
)

func sumN(n int) {
	var res = 0
	for i := 1; i <= n; i++ {
		res += i
	}
	lock.Lock() // 加锁
	factResult[n] = res
	lock.Unlock() //解锁
}

func main() {
	for i := 1; i < 200; i++ {
		go sumN(i)
	}

	time.Sleep(time.Second * 5)

	lock.Lock()
	for key, val := range factResult {
		// fmt.Printf("factorial: %v\n", factResult)
		fmt.Printf("factorial of %d is %d\n", key, val)
	}
	lock.Unlock()
}
