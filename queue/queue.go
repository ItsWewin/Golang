package queue

type Queue []int

func (queue *Queue) Push(value int) {
	*queue = append(*queue, value)
}

func (queue *Queue) Pop() int {
	tail := (*queue)[len(*queue)-1]
	*queue = (*queue)[:len(*queue)-1]

	return tail
}
