package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func firstOrder(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = append(result, node.Val)
	result = firstOrder(node.Left, result)
	result = firstOrder(node.Right, result)

	return result
}

func middleOrder(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = middleOrder(node.Left, result)
	result = append(result, node.Val)
	result = middleOrder(node.Right, result)

	return result
}

func postOrder(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = postOrder(node.Left, result)
	result = postOrder(node.Right, result)
	result = append(result, node.Val)

	return result
}

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	// write code here

	results := make([][]int, 3)

	results[0] = firstOrder(root, results[0])
	results[1] = middleOrder(root, results[1])
	results[2] = postOrder(root, results[2])

	return results

}

func main() {
	firstNode := &TreeNode{Val: 1}
	secondNode := &TreeNode{Val: 2}
	thirdNode := &TreeNode{Val: 3}
	fourthNode := &TreeNode{Val: 4}
	fifthNode := &TreeNode{Val: 5}
	sixthNode := &TreeNode{Val: 6}
	senvethNode := &TreeNode{Val: 7}

	firstNode.Left, firstNode.Right = secondNode, thirdNode
	secondNode.Left = fourthNode
	thirdNode.Left, thirdNode.Right = fifthNode, sixthNode
	fourthNode.Right = senvethNode

	results := threeOrders(firstNode)

	fmt.Println(results)
}
