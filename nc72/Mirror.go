package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类
 * @return bool布尔型
 */
func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here

	if pRoot == nil || (pRoot.Left == nil && pRoot.Right == nil) {
		return pRoot
	}

	pRoot.Left, pRoot.Right = Mirror(pRoot.Right), Mirror(pRoot.Left)

	return pRoot
}

func main() {
	// five := &TreeNode{Val: 5}
	// seven := &TreeNode{Val: 7}
	// nine := &TreeNode{Val: 9}
	// eleven := &TreeNode{Val: 11}
	// six := &TreeNode{Val: 6, Left: five, Right: seven}
	// ten := &TreeNode{Val: 10, Left: nine, Right: eleven}
	// eight := &TreeNode{Val: 8, Left: six, Right: ten}
	// Mirror(eight)

	four1 := &TreeNode{Val: 4}
	five1 := &TreeNode{Val: 5, Right: four1}
	six1 := &TreeNode{Val: 6, Right: five1}
	seven1 := &TreeNode{Val: 7, Left: six1}
	eight1 := &TreeNode{Val: 8, Left: seven1}
	Mirror(eight1)
}
