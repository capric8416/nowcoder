package main

import "fmt"

/**
 * max water
 * @param arr int整型一维数组 the array
 * @return long长整型
 */
func maxWater(arr []int) int64 {
	if len(arr) < 3 {
		return 0
	}

	result := 0

	left, right := 0, len(arr)-1
	leftMax, rightMax := arr[0], arr[right]
	for left <= right {
		if leftMax < rightMax {
			if arr[left] < leftMax {
				result += leftMax - arr[left]
			} else {
				leftMax = arr[left]
			}
			left++
		} else {
			if arr[right] < rightMax {
				result += rightMax - arr[right]
			} else {
				rightMax = arr[right]
			}
			right--
		}
	}

	return int64(result)
}

func main() {
	fmt.Println(maxWater([]int{3, 1, 2, 5, 2, 4}))
	fmt.Println(maxWater([]int{4, 5, 1, 3, 2}))
	fmt.Println(maxWater([]int{5, 7, 1, 10, 6, 2, 8, 3, 4, 9}))
}
