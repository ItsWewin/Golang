package main

import (
	"fmt"
	"learngo/queue"
)

func main() {
	myqueue := queue.Queue{}
	fmt.Println(myqueue)
	myqueue.Push(1)
	myqueue.Push(2)
	myqueue.Push(3)
	fmt.Println(myqueue)
	myqueue.Pop()
	myqueue.Pop()
	myqueue.Pop()
	// myqueue.Pop()
	fmt.Println(myqueue)
}
