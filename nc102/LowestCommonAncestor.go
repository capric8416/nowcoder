package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type (
	ChildNode struct {
		level  int
		value  *TreeNode
		parent *ChildNode
	}

	QueueNode struct {
		value *ChildNode
		next  *QueueNode
	}

	Queue struct {
		length int
		head   *QueueNode
		tail   *QueueNode
	}
)

func New() *Queue {
	return &Queue{length: 0, head: nil, tail: nil}
}

func (this *Queue) Pop() *ChildNode {
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

func (this *Queue) Push(value *ChildNode) {
	n := &QueueNode{value, nil}
	if this.length == 0 {
		this.head, this.tail = n, n
	} else {
		this.tail.next, this.tail = n, n
	}
	this.length++
}

func (this *Queue) IsEmpty() bool {
	return this.length == 0
}

/**
 *
 * @param root TreeNode类
 * @param o1 int整型
 * @param o2 int整型
 * @return int整型
 */
func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here

	var x, y *ChildNode
	q := New()
	q.Push(&ChildNode{level: 0, value: root, parent: nil})
	for !q.IsEmpty() {
		p := q.Pop()
		n := p.value
		if n.Val == o1 {
			if x == nil {
				x = p
			} else if y != nil {
				break
			}
		} else if n.Val == o2 {
			if y == nil {
				y = p
			} else if x != nil {
				break
			}
		}

		if n.Left != nil {
			q.Push(&ChildNode{level: p.level + 1, value: n.Left, parent: p})
		}
		if n.Right != nil {
			q.Push(&ChildNode{level: p.level + 1, value: n.Right, parent: p})
		}
	}

	if y.level > x.level {
		x, y = y, x
	}

	for x.level > y.level {
		x = x.parent
	}
	if x == y {
		return x.value.Val
	}

	for i := x.level; i >= 0; {
		if x.parent == y.parent {
			return x.parent.value.Val
		} else {
			x, y = x.parent, y.parent
		}
	}

	return -1
}


func main() {
	seven := &TreeNode{Val: 7}
	four := &TreeNode{Val: 4}
	zero := &TreeNode{Val: 0}
	eight := &TreeNode{Val: 8}
	six := &TreeNode{Val: 6}
	two := &TreeNode{Val: 2, Left: seven, Right: four}
	five := &TreeNode{Val: 5, Left: six, Right: two}
	one := &TreeNode{Val: 1, Left: zero, Right: eight}
	three := &TreeNode{Val: 3, Left: five, Right: one}
	ten := &TreeNode{Val: 10, Left: three}
	fmt.Println(lowestCommonAncestor(ten, 5, 1))
	fmt.Println(lowestCommonAncestor(ten, 7, 1))
	fmt.Println(lowestCommonAncestor(ten, 4, 8))
	fmt.Println(lowestCommonAncestor(ten, 5, 4))
	fmt.Println(lowestCommonAncestor(ten, 6, 4))
	fmt.Println(lowestCommonAncestor(ten, 6, 7))
	fmt.Println(lowestCommonAncestor(ten, 8, 0))

	CreateBinTreeFromString("27,
	32,34,
	19,41,17,18,
	9,14,44,39,#,#,24,30,
	#,#,#,2,7,42,28,36,#,#,11,6,#,1,#,#,#,31,16,4,22,33,#,#,#,5,10,15,37,12,8,#,35,3,#,23,21,#,#,#,29,#,#,#,40,#,#,#,#,#,#,#,#,#,13,43,#,#,#,#,#,#,25,20,#,#,38,#,26")
}
