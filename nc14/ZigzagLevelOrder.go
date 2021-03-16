package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return [][]int{}
	}

	results := make([][]int, 0)

	nodes := []*TreeNode{root}
	count := len(nodes)
	for i, reverse := 0, false; count > 0; i, count, reverse = i+1, len(nodes), !reverse {
		results = append(results, make([]int, count))

		if !reverse {
			for j := 0; j < count; j++ {
				n := nodes[j]
				results[i][j] = n.Val
			}
		} else {
			for j := 0; j < count; j++ {
				n := nodes[j]
				results[i][count-j-1] = n.Val
			}
		}

		for _, n := range nodes {
			if n.Left != nil {
				nodes = append(nodes, n.Left)
			}
			if n.Right != nil {
				nodes = append(nodes, n.Right)
			}
		}
		nodes = nodes[count:]
	}

	return results
}

func main() {
	fifteen := &TreeNode{Val: 15}
	seven := &TreeNode{Val: 7}
	twenty := &TreeNode{Val: 20, Left: fifteen, Right: seven}
	nine := &TreeNode{Val: 9}
	three := &TreeNode{Val: 3, Left: nine, Right: twenty}

	fmt.Println(zigzagLevelOrder(three))

	four := &TreeNode{Val: 4}
	two := &TreeNode{Val: 2, Left: four}
	five := &TreeNode{Val: 5}
	three = &TreeNode{Val: 3, Right: five}
	one := &TreeNode{Val: 1, Left: two, Right: three}

	fmt.Println(zigzagLevelOrder(one))
}
