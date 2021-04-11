package main

import "fmt"

func dfs(grid [][]byte, x, y int, m, n int) {
	if x < 0 || x == m || y < 0 || y == n || grid[x][y] == '0' {
		return
	}

	grid[x][y] = '0'
	dfs(grid, x-1, y, m, n)
	dfs(grid, x+1, y, m, n)
	dfs(grid, x, y-1, m, n)
	dfs(grid, x, y+1, m, n)
}

/**
 * 判断岛屿数量
 * @param grid char字符型二维数组
 * @return int整型
 */
func solve(grid [][]byte) int {
	// write code here

	result := 0

	m, n := len(grid), len(grid[0])
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == '1' {
				result++
				dfs(grid, x, y, m, n)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(solve([][]byte{{'1', '1', '0', '0', '0'}, {'0', '1', '0', '1', '1'}, {'0', '0', '0', '1', '1'}, {'0', '0', '0', '0', '0'}, {'0', '0', '1', '1', '1'}}))
	fmt.Println(solve([][]byte{{'1', '1', '0', '0', '0'}, {'0', '1', '1', '1', '1'}, {'1', '1', '0', '1', '1'}, {'0', '0', '0', '0', '0'}, {'0', '0', '1', '1', '1'}}))
}
