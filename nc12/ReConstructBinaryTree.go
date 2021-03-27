package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rebuild(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	n := &TreeNode{Val: pre[0]}

	var i int
	for i = 0; i < len(vin); i++ {
		if vin[i] == pre[0] {
			break
		}
	}

	n.Left = rebuild(pre[1:i+1], vin[:i])
	n.Right = rebuild(pre[i+1:], vin[i+1:])

	return n
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here

	return rebuild(pre, vin)
}

func main() {
	t1 := reConstructBinaryTree([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 2, 4, 1, 6, 5, 7})
	t1 = reConstructBinaryTree([]int{1, 2, 4, 6, 7, 8, 3, 5}, []int{4, 7, 6, 8, 2, 1, 3, 5})
	t1 = reConstructBinaryTree([]int{1, 2, 4, 5, 7, 8, 3, 6}, []int{4, 2, 7, 5, 8, 1, 3, 6})
	_ = t1
}
