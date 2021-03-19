package main

/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here

	half := len(numbers) / 2
	count := make(map[int]int, 0)
	for _, n := range numbers {
		count[n]++
		if count[n] > half {
			return n
		}
	}

	return 0
}
