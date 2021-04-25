package main

import (
	"strconv"
	"strings"
)

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

func serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	queue := NewQueue()
	queue.Push(root)

	nodes := make([]string, 0)
	for i := 0; queue.Len() > 0; i++ {
		count := queue.Len()
		for j := 0; j < count; j++ {
			current := queue.Pop()
			if current == nil {
				nodes = append(nodes, "#")
			} else {
				nodes = append(nodes, strconv.Itoa(current.Val))
				queue.Push(current.Left)
				queue.Push(current.Right)
			}
		}
	}

	return strings.Join(nodes, ",")
}

func deserialize(text string) *TreeNode {
	if len(text) == 0 {
		return nil
	}

	queue := NewQueue()

	nodes := strings.Split(text, ",")
	v, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: v}
	queue.Push(root)

	i := 0
	for queue.Len() > 0 {
		count := queue.Len()
		for j := 0; j < count; j++ {
			current := queue.Pop()
			if current == nil {
				continue
			}

			i++
			v, err := strconv.Atoi(nodes[i])
			if err == nil {
				current.Left = &TreeNode{Val: v}
			}

			i++
			v, err = strconv.Atoi(nodes[i])
			if err == nil {
				current.Right = &TreeNode{Val: v}
			}

			queue.Push(current.Left)
			queue.Push(current.Right)
		}
	}

	return root
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
 */
func Serialize(root *TreeNode) *TreeNode {
	// write code here

	text := serialize(root)

	node := deserialize(text)

	return node
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2, Right: one}
	three := &TreeNode{Val: 3, Left: two}
	six := &TreeNode{Val: 6}
	eight := &TreeNode{Val: 8}
	seven := &TreeNode{Val: 7, Left: six, Right: eight}
	five := &TreeNode{Val: 5, Left: three, Right: seven}
	Serialize(five)
}
