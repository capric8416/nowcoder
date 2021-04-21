package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, value int, sum int, path []int, paths *[][]int) {
	path = append(path, root.Val)

	value += root.Val

	if root.Left == nil && root.Right == nil {
		if value == sum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			*paths = append(*paths, tmp)
		}
		return
	}

	if root.Left != nil {
		dfs(root.Left, value, sum, path, paths)
	}

	if root.Right != nil {
		dfs(root.Right, value, sum, path, paths)
	}
}

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
func pathSum(root *TreeNode, sum int) [][]int {
	// write code here
	if root == nil {
		return [][]int{}
	}

	result := make([]int, 0)
	results := make([][]int, 0)
	dfs(root, 0, sum, result, &results)

	return results
}

func main() {
	eight4 := &TreeNode{Val: 1}
	seven4 := &TreeNode{Val: 1}
	six4 := &TreeNode{Val: 1}
	five4 := &TreeNode{Val: 1}
	four4 := &TreeNode{Val: 1}
	three4 := &TreeNode{Val: 1}
	two4 := &TreeNode{Val: 1}
	one4 := &TreeNode{Val: 1}
	four3 := &TreeNode{Val: 1, Left: seven4, Right: eight4}
	three3 := &TreeNode{Val: 1, Left: five4, Right: six4}
	two3 := &TreeNode{Val: 1, Left: three4, Right: four4}
	one3 := &TreeNode{Val: 1, Left: one4, Right: two4}
	two2 := &TreeNode{Val: 2, Left: three3, Right: four3}
	one2 := &TreeNode{Val: 1, Left: one3, Right: two3}
	one1 := &TreeNode{Val: 1, Left: one2, Right: two2}
	fmt.Println(pathSum(one1, 4))
}
