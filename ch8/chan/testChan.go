package testChan

import "time"

func testChanEqual(ch1, ch2 chan int) bool {
	return ch1 == ch2
}

func testChanReceive(ch chan int) int {
	result := <-ch
	time.Sleep(1 * time.Second)
	return result
}
