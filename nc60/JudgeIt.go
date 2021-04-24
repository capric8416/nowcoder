package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func gt(a int, b *int) bool {
	return b != nil && a > *b
}

func lt(a int, b *int) bool {
	return b != nil && a < *b
}

func dfs(root *TreeNode, depth int, maxDepth, secondDepth *int, ancestor *int, compare func(int, *int) bool, search, complete *bool) {
	if *search == false && *complete == false {
		return
	}

	depth += 1

	if *complete == true && (root.Left == nil || root.Right == nil) {
		if root.Right != nil && root.Left == nil {
			*complete = false
		} else {
			if *maxDepth == 0 {
				*maxDepth, *secondDepth = depth, depth
			} else {
				maxSub, secondSub := *maxDepth-depth, *secondDepth-depth
				if maxSub == 0 {
				} else if maxSub == 1 {
					*secondDepth = depth
				} else if maxSub < 0 || maxSub > 1 || secondSub < 0 || secondSub > 0 {
					*complete = false
				}
			}
		}
	}

	if root.Left != nil {
		if (root.Left.Val > root.Val) || compare(root.Left.Val, ancestor) {
			*search = false
		}
	}

	if root.Right != nil {
		if (root.Right.Val < root.Val) || compare(root.Right.Val, ancestor) {
			*search = false
		}
	}

	if root.Left != nil {
		dfs(root.Left, depth, maxDepth, secondDepth, &root.Val, gt, search, complete)
	}

	if root.Right != nil {
		dfs(root.Right, depth, maxDepth, secondDepth, &root.Val, lt, search, complete)
	}
}

/**
 *
 * @param root TreeNode类 the root
 * @return bool布尔型一维数组
 */
func judgeIt(root *TreeNode) []bool {
	// write code here

	if root == nil {
		return []bool{true, true}
	}

	search, complete, maxDepth, secondDepth := true, true, 0, 0
	dfs(root, 0, &maxDepth, &secondDepth, nil, gt, &search, &complete)
	return []bool{search, complete}
}

func main() {
	// one := &TreeNode{Val: 1}
	// three := &TreeNode{Val: 3}
	// two := &TreeNode{Val: 2, Left: one, Right: three}
	// five := &TreeNode{Val: 5}
	// seven := &TreeNode{Val: 7}
	// six := &TreeNode{Val: 6, Left: five, Right: seven}
	// four := &TreeNode{Val: 4, Left: two, Right: six}
	// nine := &TreeNode{Val: 9}
	// ten := &TreeNode{Val: 10, Left: nine}
	// // twelve := &TreeNode{Val: 12}
	// eleven := &TreeNode{Val: 11, Left: ten /*, Right: twelve*/}
	// eight := &TreeNode{Val: 8, Left: four, Right: eleven}
	// fmt.Println(judgeIt(eight))

	one := &TreeNode{Val: 1}
	three := &TreeNode{Val: 3}
	two := &TreeNode{Val: 2, Left: one, Right: three}
	five := &TreeNode{Val: 5}
	seven := &TreeNode{Val: 7}
	six := &TreeNode{Val: 6, Left: five, Right: seven}
	four := &TreeNode{Val: 4, Left: two, Right: six}
	// nine := &TreeNode{Val: 9}
	// fourteen := &TreeNode{Val: 14}
	thirdteen := &TreeNode{Val: 13 /*, Left: nine , Right: fourteen*/}
	// seventeen := &TreeNode{Val: 17}
	// nineteen := &TreeNode{Val: 19}
	// egihteen := &TreeNode{Val: 18 , Left: seventeen, Right: nineteen}
	sixteen := &TreeNode{Val: 16, Left: thirdteen /*, Right: egihteen*/}
	eight := &TreeNode{Val: 8, Left: four, Right: sixteen}
	fmt.Println(judgeIt(eight))
}
