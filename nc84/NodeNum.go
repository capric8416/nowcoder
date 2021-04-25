package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pow(base, count int) int {
	result := 1
	for i := 0; i < count; i++ {
		result *= base
	}
	return result
}

func getHeight(head *TreeNode) int {
	height := 0
	for curr := head; curr != nil; curr = curr.Left {
		height++
	}
	return height
}

/**
 *
 * @param head TreeNode类
 * @return int整型
 */
func nodeNum(head *TreeNode) int {
	// write code here

	if head == nil {
		return 0
	}

	leftHeight := getHeight(head.Left)
	rightHeight := getHeight(head.Right)
	if leftHeight == rightHeight {
		return pow(2, leftHeight) + nodeNum(head.Right)
	}
	return pow(2, rightHeight) + nodeNum(head.Left)
}

func main() {
	one := &TreeNode{Val: 1}
	fmt.Println(nodeNum(one))

	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	one = &TreeNode{Val: 1, Left: two, Right: three}
	fmt.Println(nodeNum(one))
}
