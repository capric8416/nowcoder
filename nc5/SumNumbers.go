package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, value int, sum *int) {
	value = value*10 + root.Val

	if root.Left == nil && root.Right == nil {
		*sum += value
		return
	}

	if root.Left != nil {
		dfs(root.Left, value, sum)
	}

	if root.Right != nil {
		dfs(root.Right, value, sum)
	}
}

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
func sumNumbers(root *TreeNode) int {
	// write code here
	sum := 0
	if root != nil {
		dfs(root, 0, &sum)
	}
	return sum
}

func main() {
	eight4 := &TreeNode{Val: 8}
	seven4 := &TreeNode{Val: 7}
	six4 := &TreeNode{Val: 6}
	five4 := &TreeNode{Val: 5}
	four4 := &TreeNode{Val: 4}
	three4 := &TreeNode{Val: 3}
	two4 := &TreeNode{Val: 2}
	one4 := &TreeNode{Val: 1}
	four3 := &TreeNode{Val: 4, Left: seven4, Right: eight4}
	three3 := &TreeNode{Val: 3, Left: five4, Right: six4}
	two3 := &TreeNode{Val: 2, Left: three4, Right: four4}
	one3 := &TreeNode{Val: 1, Left: one4, Right: two4}
	two2 := &TreeNode{Val: 2, Left: three3, Right: four3}
	one2 := &TreeNode{Val: 1, Left: one3, Right: two3}
	one1 := &TreeNode{Val: 1, Left: one2, Right: two2}
	fmt.Println(sumNumbers(one1))
}
