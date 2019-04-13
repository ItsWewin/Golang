package gotest2

func add(a, b int) int {
	return a + b
}

func accumulate(num int) int {
	if num == 1 {
		return num
	}

	return num + accumulate(num-1)
}
