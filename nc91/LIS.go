package main

import "fmt"

func UpperBound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left < right {
		center := left + (right-left)>>1
		if arr[center] < target {
			left = center + 1
		} else {
			right = center
		}
	}

	return left
}

/**
 * retrun the longest increasing subsequence
 * @param arr int整型一维数组 the array
 * @return int整型一维数组
 */
func LIS(arr []int) []int {
	// write code here

	length := len(arr)
	if length < 2 {
		return arr
	}

	result := []int{arr[0]}
	maxLength := []int{1}
	for i, right := 1, 0; i < length; i++ {
		v := arr[i]
		if v > result[right] {
			result = append(result, v)
			maxLength = append(maxLength, len(result))
			right++
		} else {
			pos := UpperBound(result, v)
			result[pos] = v
			maxLength = append(maxLength, pos+1)
		}
	}

	for i, j := length-1, len(result); j > 0; i-- {
		if maxLength[i] == j {
			j--
			if result[j] != arr[i] {
				result[j] = arr[i]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(LIS([]int{}))
	fmt.Println(LIS([]int{1}))
	fmt.Println(LIS([]int{2, 1, 5, 3, 6, 4, 8, 9, 7}))
	fmt.Println(LIS([]int{1, 2, 8, 6, 4}))
	fmt.Println(LIS([]int{2, 3, 1, 2, 3}))
}
