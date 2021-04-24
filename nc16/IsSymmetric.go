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
func New() *Queue {
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
 * @return bool布尔型
 */
func isSymmetric(root *TreeNode) bool {
	// write code here

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	queue := New()
	queue.Push(root.Left)
	queue.Push(root.Right)

	for queue.Len() > 0 {
		left := queue.Pop()
		right := queue.Pop()

		// fmt.Printf("left: %v, right: %v\n", left, right)

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue.Push(left.Left)
		queue.Push(right.Right)
		queue.Push(left.Right)
		queue.Push(right.Left)
	}

	return true
}

func main() {
	five1 := &TreeNode{Val: 5}
	five2 := &TreeNode{Val: 5}
	four1 := &TreeNode{Val: 4, Right: five1}
	four2 := &TreeNode{Val: 4}
	four3 := &TreeNode{Val: 4}
	four4 := &TreeNode{Val: 4}
	four5 := &TreeNode{Val: 4}
	four6 := &TreeNode{Val: 4}
	four7 := &TreeNode{Val: 4}
	four8 := &TreeNode{Val: 4, Right: five2}
	three1 := &TreeNode{Val: 3, Left: four1, Right: four2}
	three2 := &TreeNode{Val: 3, Left: four3, Right: four4}
	three3 := &TreeNode{Val: 3, Left: four5, Right: four6}
	three4 := &TreeNode{Val: 3, Left: four7, Right: four8}
	two1 := &TreeNode{Val: 2, Left: three1, Right: three2}
	two2 := &TreeNode{Val: 2, Left: three3, Right: three4}
	one := &TreeNode{Val: 1, Left: two1, Right: two2}
	fmt.Println(isSymmetric(one))
}
