package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := TreeDepth(root.Left)
	if left == -1 {
		return -1
	}

	right := TreeDepth(root.Right)
	if right == -1 {
		return -1
	}

	sub := left - right
	if sub < -1 || sub > 1 {
		return -1
	}

	if sub >= 0 {
		return left + 1
	} else {
		return right + 1
	}
}

/**
 *
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here

	return TreeDepth(pRoot) != -1
}
