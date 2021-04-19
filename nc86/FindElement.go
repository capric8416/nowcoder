package main

import "fmt"

/**
 *
 * @param mat int整型二维数组
 * @param n int整型
 * @param m int整型
 * @param x int整型
 * @return int整型一维数组
 */
func findElement(mat [][]int, n int, m int, x int) []int {
	// write code here

	for row := 0; row < n; row++ {
		if mat[row][0] > x || mat[row][m-1] < x {
			continue
		}
		for left, right := 0, m-1; left <= right; {
			center := left + (right-left)>>1
			if mat[row][center] == x {
				return []int{row, center}
			} else if mat[row][center] < x {
				left = center + 1
			} else {
				right = center - 1
			}
		}
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(findElement([][]int{{1, 2, 3}, {4, 5, 6}}, 2, 3, 6))
}
