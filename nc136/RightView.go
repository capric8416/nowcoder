package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rebuild(xianxu []int, zhongxu []int) *TreeNode {
	if len(xianxu) == 0 {
		return nil
	}

	root := &TreeNode{Val: xianxu[0]}

	var i int
	for i = 0; i < len(zhongxu); i++ {
		if zhongxu[i] == xianxu[0] {
			break
		}
	}
	root.Left = rebuild(xianxu[1:i+1], zhongxu[:i])
	root.Right = rebuild(xianxu[i+1:], zhongxu[i+1:])

	return root
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
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 求二叉树的右视图
 * @param xianxu int整型一维数组 先序遍历
 * @param zhongxu int整型一维数组 中序遍历
 * @return int整型一维数组
 */
func solve(xianxu []int, zhongxu []int) []int {
	// write code here

	root := rebuild(xianxu, zhongxu)

	queue := NewQueue()
	queue.Push(root)

	result := make([]int, 0)
	for i := 0; queue.Len() > 0; i++ {
		count := queue.Len()
		for j := 0; j < count; j++ {
			current := queue.Pop()
			if j == count-1 {
				result = append(result, current.Val)
			}

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
	fmt.Println(solve([]int{1, 2, 4, 5, 3}, []int{4, 2, 5, 1, 3}))
}
