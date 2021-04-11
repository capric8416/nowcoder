package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func middleOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	middleOrder(node.Left, result)
	*result = append(*result, node.Val)
	middleOrder(node.Right, result)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */
func inorderTraversal(root *TreeNode) []int {
	// write code here

	result := make([]int, 0)
	middleOrder(root, &result)
	return result
}

func main() {
	second := &TreeNode{Val: 2}
	third := &TreeNode{Val: 3}
	first := &TreeNode{Val: 1, Left: second, Right: third}
	fmt.Println(inorderTraversal(first))
}
