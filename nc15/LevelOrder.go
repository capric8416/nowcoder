package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type (
	Queue struct {
		head, tail *QueueNode
		length     int
	}
	QueueNode struct {
		value *TreeNode
		next  *QueueNode
	}
)

// Create a new queue
func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil, length: 0}
}

// Take the next item off the front of the queue
func (this *Queue) Pop() *TreeNode {
	if this.length == 0 {
		return nil
	}
	n := this.head
	if this.length == 1 {
		this.head, this.tail = nil, nil
	} else {
		this.head = this.head.next
	}
	this.length--
	return n.value
}

// Put an item on the end of a queue
func (this *Queue) Push(value *TreeNode) {
	n := &QueueNode{value, nil}
	if this.length == 0 {
		this.head, this.tail = n, n
	} else {
		this.tail.next, this.tail = n, n
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
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	// write code here

	if root == nil {
		return [][]int{}
	}

	s := NewQueue()
	s.Push(root)

	values := make([][]int, 0)
	for i := 0; s.Len() > 0; i++ {
		values = append(values, make([]int, 0))

		count := s.Len()
		for j := 0; j < count; j++ {
			current := s.Pop()
			values[i] = append(values[i], current.Val)

			if current.Left != nil {
				s.Push(current.Left)
			}
			if current.Right != nil {
				s.Push(current.Right)
			}
		}
	}

	return values
}

func main() {
	firstNode := &TreeNode{Val: 1}
	secondNode := &TreeNode{Val: 2}
	thirdNode := &TreeNode{Val: 3}
	fourthNode := &TreeNode{Val: 4}
	fifthNode := &TreeNode{Val: 5}
	sixthNode := &TreeNode{Val: 6}
	senvethNode := &TreeNode{Val: 7}

	firstNode.Left, firstNode.Right = secondNode, thirdNode
	secondNode.Left = fourthNode
	thirdNode.Left, thirdNode.Right = fifthNode, sixthNode
	fourthNode.Right = senvethNode

	result := levelOrder(firstNode)
	fmt.Println(result)
}
