package main

import "fmt"

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	// write code here

	cache := make(map[int]int, 0)
	cache[numbers[0]] = 0

	length := len(numbers)
	for j := 1; j < length; j++ {
		sub := target - numbers[j]
		i, ok := cache[sub]
		if ok {
			return []int{i + 1, j + 1}
		}

		cache[numbers[j]] = j
	}

	return []int{}
}

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)

	result = twoSum([]int{20, 70, 110, 150}, 90)
	fmt.Println(result)
}
