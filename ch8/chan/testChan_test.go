package testChan

import (
	"fmt"
	"testing"
)

func Test_TestChan(t *testing.T) {
	ch1 := make(chan int)
	// ch2 := make(chan int)
	ch2 := ch1
	if result := testChanEqual(ch1, ch2); result {
		fmt.Printf("successd!")
	} else {
		t.Errorf("testChanEqual() get %t, expected %t", result, !result)
	}
}

func Test_testChanSend(t *testing.T) {
	ch := make(chan int)
	ch <- 12
	if result := testChanReceive(ch); result == 12 {
		fmt.Print("succeed")
	} else {
		t.Errorf("testChanReceive() get %d, expected %d", result, 12)
	}
}
