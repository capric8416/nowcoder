package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, value int, sum int, result *bool) {
	if *result == true {
		return
	}

	value += root.Val

	if root.Left == nil && root.Right == nil {
		if value == sum {
			*result = true
		}
		return
	}

	if root.Left != nil {
		dfs(root.Left, value, sum, result)
	}

	if root.Right != nil {
		dfs(root.Right, value, sum, result)
	}
}

/**
 *
 * @param root TreeNode类
 * @param sum int整型
 * @return bool布尔型
 */
func hasPathSum(root *TreeNode, sum int) bool {
	// write code here

	if root == nil {
		return false
	}

	result := false
	dfs(root, 0, sum, &result)
	return result
}
