package main

import "fmt"

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
 * @param root TreeNode类
 * @return int整型
 */
func maxDepth(root *TreeNode) int {
	// write code here

	if root == nil {
		return 0
	}

	queue := NewQueue()
	queue.Push(root)

	result := 0
	for i := 0; queue.Len() > 0; i++ {
		result++

		count := queue.Len()
		for j := 0; j < count; j++ {
			current := queue.Pop()
			if current.Left != nil {
				queue.Push(current.Left)
			}
			if current.Right != nil {
				queue.Push(current.Right)
			}
		}
	}

	return result
}

func main() {
	second := &TreeNode{Val: 2}
	first := &TreeNode{Val: 1, Left: second}
	fmt.Println(maxDepth(first))

	four := &TreeNode{Val: 4}
	second = &TreeNode{Val: 2, Left: four}
	five := &TreeNode{Val: 5}
	three := &TreeNode{Val: 3, Right: five}
	first = &TreeNode{Val: 1, Left: second, Right: three}
	fmt.Println(maxDepth(first))
}
