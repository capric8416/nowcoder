package main

import "fmt"

/**
 *
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxLength(arr []int) int {
	// write code here

	if len(arr) < 2 {
		return len(arr)
	}

	maxSubLength := 1

	valueIndex := map[int]int{arr[0]: 0}
	for left, right := 0, 1; right < len(arr); right++ {
		value := arr[right]
		// 正常操作，这里应该删除valueIndex中从left到index之间的值
		// 但这里通过比较index与left的大小，亦可判断命中区间是否有效
		if index, exists := valueIndex[value]; exists && index >= left {
			left = index + 1
		}

		valueIndex[value] = right

		subLength := right - left + 1
		if subLength > maxSubLength {
			maxSubLength = subLength
		}
	}

	return maxSubLength
}

func main() {
	fmt.Println(maxLength([]int{}))
	fmt.Println(maxLength([]int{2}))
	fmt.Println(maxLength([]int{2, 3, 4, 5}))
	fmt.Println(maxLength([]int{2, 2, 3, 4, 3}))
	fmt.Println(maxLength([]int{2, 2, 3, 4, 5, 6, 7, 4}))
}
