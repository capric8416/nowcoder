package main

import "fmt"

/**
 * 旋转数组
 * @param n int整型 数组长度
 * @param m int整型 右移距离
 * @param a int整型一维数组 给定数组
 * @return int整型一维数组
 */
func solve(n int, m int, a []int) []int {
	// write code here

	if m == 0 || len(a) == 1 {
		return a
	}

	if m > n {
		m %= n
	}

	if m == n {
		return a
	}
	m = n - m

	for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for i, j := m, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	fmt.Println("->", []int{1, 2, 3, 4, 5, 6})
	fmt.Println("<-", solve(6, 2, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println("<-", solve(6, 6, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println("<-", solve(6, 7, []int{1, 2, 3, 4, 5, 6}))
}
