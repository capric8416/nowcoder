package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode

func inOrder(curr *TreeNode, result []int, init *bool) {
	if curr == nil {
		return
	}

	inOrder(curr.Left, result, init)

	if prev != nil && prev.Val > curr.Val && *init == true {
		result[0] = curr.Val
	}
	if prev != nil && prev.Val > curr.Val && *init == false {
		result[1] = prev.Val
		*init = true
	}

	prev = curr

	inOrder(curr.Right, result, init)
}

/**
 *
 * @param root TreeNode类 the root
 * @return int整型一维数组
 */
func findError(root *TreeNode) []int {
	// write code here

	init, result := false, make([]int, 2)
	inOrder(root, result, &init)
	return result
}

func main() {
	one := &TreeNode{Val: 1}
	five := &TreeNode{Val: 5}
	two := &TreeNode{Val: 2, Left: one, Right: five}
	three := &TreeNode{Val: 3}
	seven := &TreeNode{Val: 7}
	six := &TreeNode{Val: 6, Left: three, Right: seven}
	four := &TreeNode{Val: 4, Left: two, Right: six}
	fmt.Println(findError(four))
}
