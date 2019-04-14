package gotest

func add(a, b int) int {
	return a + b
}

// 递归计算前 num 个正整数的和
// 为什么没有在 doc 中显示出来
func accumulate(num int) int {
	if num == 1 {
		return num
	}

	return num + accumulate(num-1)
}
