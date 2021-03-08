package main

import "fmt"

func minHeapSink(arr []int, parentIndex int) []int {
	length := len(arr)

	for childIndex := 2*parentIndex + 1; childIndex < length; childIndex = 2*parentIndex + 1 {
		if childIndex+1 < length && arr[childIndex+1] < arr[childIndex] {
			childIndex++
		}

		if arr[parentIndex] <= arr[childIndex] {
			break
		}

		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]

		parentIndex = childIndex
	}

	return arr
}

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	// write code here

	if n == 0 || K == 0 || K > n {
		return -1
	}

	result := a[:K]

	for i := (K - 1) / 2; i >= 0; i-- {
		result = minHeapSink(result, i)
	}

	for i := K; i < n; i++ {
		if a[i] > result[0] {
			result[0] = a[i]
			result = minHeapSink(result, 0)
		}
	}

	return result[0]
}

func main() {
	n := findKth([]int{1332802, 1177178, 1514891, 871248, 753214, 123866, 1615405, 328656, 1540395, 968891, 1884022, 252932, 1034406, 1455178, 821713, 486232, 860175, 1896237, 852300, 566715, 1285209, 1845742, 883142, 259266, 520911, 1844960, 218188, 1528217, 332380, 261485, 1111670, 16920, 1249664, 1199799, 1959818, 1546744, 1904944, 51047, 1176397, 190970, 48715, 349690, 673887, 1648782, 1010556, 1165786, 937247, 986578, 798663}, 49, 24)
	fmt.Println(n)
}
