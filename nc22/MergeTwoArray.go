package main

import "fmt"

/**
 *
 * @param A int整型一维数组
 * @param B int整型一维数组
 * @return void
 */
func merge(A []int, m int, B []int, n int) {
	// write code here

	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; k-- {
		if A[i] < B[j] {
			A[k] = B[j]
			j--
		} else {
			A[k] = A[i]
			i--
		}
	}

	for ; j >= 0; j, k = j-1, k-1 {
		A[k] = B[j]
	}
}

func main() {
	A := make([]int, 10)
	copy(A, []int{1, 2, 4, 6})
	merge(A, 4, []int{0, 1, 3, 5, 7, 8}, 6)
	fmt.Println(A)

	A = make([]int, 1)
	merge(A, 0, []int{1}, 1)
	fmt.Println(A)

	A = make([]int, 1)
	copy(A, []int{1})
	merge(A, 1, []int{}, 0)
	fmt.Println(A)
}
