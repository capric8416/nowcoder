package main

import "fmt"

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here

	if number == 0 || number == 1 {
		return number
	}

	current, prev2, prev1 := 0, 1, 1
	for i := 2; i <= number; i++ {
		current = prev2 + prev1
		prev2 = prev1
		prev1 = current
	}

	return current
}

func main() {
	for i := 0; i < 30; i++ {
		count := jumpFloor(i)
		fmt.Println(i, "->", count)
	}
}
