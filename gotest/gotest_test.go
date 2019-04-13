package gotest2

import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{2, 3, 5},
	}

	for _, v := range tests {
		if sum := add(v.a, v.b); sum == v.c {
			t.Log("OK")
		} else {
			t.Errorf("add(%d, %d); got %d, expected %d", v.a, v.b, sum, v.c)
		}
	}
}

func Test_Accumulate(t *testing.T) {
	tests := []struct {
		a, b int
	}{
		{3, 6},
		{4, 10},
		{5, 15},
	}

	for _, v := range tests {
		if sum := accumulate(v.a); sum == v.b {
			fmt.Println("OK")
		} else {
			t.Errorf("accumulate(%d); got %d, expected %d", v.a, sum, v.b)
			// t.Error("测试失败！")
		}
	}
}
