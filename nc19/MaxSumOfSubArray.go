package main

import "fmt"

/**
 * max sum of the subarray
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxsumofSubarray(arr []int) int {
	// write code here

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		sum := arr[i-1] + arr[i]
		if sum > arr[i] {
			arr[i] = sum
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func main() {
	result := maxsumofSubarray([]int{1, -2, 3, 5, -2, 6, -1})
	fmt.Println(result)
}
