package main

import "fmt"

/**
 *
 * @param matrix int整型二维数组 the matrix
 * @return int整型
 */
func minPathSum(matrix [][]int) int {
	// write code here

	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	for x := 1; x < m; x++ {
		matrix[x][0] += matrix[x-1][0]
	}
	for y := 1; y < n; y++ {
		matrix[0][y] += matrix[0][y-1]
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			if matrix[x-1][y] <= matrix[x][y-1] {
				matrix[x][y] += matrix[x-1][y]
			} else {
				matrix[x][y] += matrix[x][y-1]
			}
		}
	}

	return matrix[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 5, 9}, {8, 1, 3, 4}, {5, 0, 6, 1}, {8, 8, 4, 0}}))
}
