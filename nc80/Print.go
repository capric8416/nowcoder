package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type (
	Queue struct {
		start, end *QueueNode
		length     int
	}
	QueueNode struct {
		value *TreeNode
		next  *QueueNode
	}
)

// Create a new queue
func NewQueue() *Queue {
	return &Queue{start: nil, end: nil, length: 0}
}

// Take the next item off the front of the queue
func (this *Queue) Pop() *TreeNode {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return n.value
}

// Put an item on the end of a queue
func (this *Queue) Push(value *TreeNode) {
	n := &QueueNode{value, nil}
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}

// Return the number of items in the queue
func (this *Queue) Len() int {
	return this.length
}

/**
 *
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	// write code here

	if pRoot == nil {
		return [][]int{}
	}

	fifo := NewQueue()
	fifo.Push(pRoot)

	values := make([][]int, 0)
	for i := 0; fifo.Len() > 0; i++ {
		values = append(values, make([]int, 0))

		count := fifo.Len()
		for j := 0; j < count; j++ {
			current := fifo.Pop()
			values[i] = append(values[i], current.Val)

			if current.Left != nil {
				fifo.Push(current.Left)
			}
			if current.Right != nil {
				fifo.Push(current.Right)
			}
		}
	}

	return values
}
