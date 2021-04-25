package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param num int整型一维数组
 * @return TreeNode类
 */
func sortedArrayToBST(num []int) *TreeNode {
	// write code here

	length := len(num)
	if length == 0 {
		return nil
	}

	center := length >> 1
	return &TreeNode{Val: center, Left: sortedArrayToBST(num[:center]), Right: sortedArrayToBST(num[center+1:])}
}
