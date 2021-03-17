package main

import "fmt"

/**
 *
 * @param matrix int整型二维数组
 * @return int整型一维数组
 */
func spiralOrder(matrix [][]int) []int {
	// write code here

	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}

	columns := len(matrix[0])

	count := rows * columns
	result := make([]int, count, count)

	for i, r, c, row, col := 0, 0, 0, rows-1, columns-1; r <= row && c <= col; {
		for cc := c; cc <= col; cc, i = cc+1, i+1 {
			result[i] = matrix[r][cc]
		}
		r++

		for rr := r; rr <= row; rr, i = rr+1, i+1 {
			result[i] = matrix[rr][col]
		}
		col--

		if r <= row {
			for cc := col; cc >= c; cc, i = cc-1, i+1 {
				result[i] = matrix[row][cc]
			}
		}
		row--

		if c <= col {
			for rr := row; rr >= r; rr, i = rr-1, i+1 {
				result[i] = matrix[rr][c]
			}
		}
		c++
	}

	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{}))

	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}}))

	fmt.Println(spiralOrder([][]int{{1}, {2}, {3}, {4}}))

	matrix := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
		{101, 102, 103, 104, 105, 106, 107, 108, 109, 110},
	}
	fmt.Println(spiralOrder(matrix))
}
