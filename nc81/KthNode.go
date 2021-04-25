package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result *TreeNode

func inOrder(root *TreeNode, k *int) {
	if root == nil || result != nil {
		return
	}

	inOrder(root.Left, k)

	(*k)--
	if *k == 0 {
		result = root
		return
	}

	inOrder(root.Right, k)
}

/**
 *
 * @param pRoot TreeNode类
 * @param k int整型
 * @return TreeNode类
 */
func KthNode(pRoot *TreeNode, k int) *TreeNode {
	// write code here

	result = nil
	inOrder(pRoot, &k)
	return result
}

func main() {
	two := &TreeNode{Val: 2}
	four := &TreeNode{Val: 4}
	three := &TreeNode{Val: 3, Left: two, Right: four}
	six := &TreeNode{Val: 6}
	eight := &TreeNode{Val: 8}
	seven := &TreeNode{Val: 7, Left: six, Right: eight}
	five := &TreeNode{Val: 5, Left: three, Right: seven}
	for i := 1; i < 8; i++ {
		fmt.Println(KthNode(five, i))
	}
}
