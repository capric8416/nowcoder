package main

import "fmt"

/**
 *
 * @param x int整型
 * @return int整型
 */
func sqrt(x int) int {
	// write code here
	if x < 1 {
		return 0
	}
	if x < 4 {
		return 1
	}

	left, right := 2, x
	for {
		center := (right-left)>>1 + left
		if center*center <= x && (center+1)*(center+1) > x {
			return center
		} else if center*center < x {
			left = center + 1
		} else {
			right = center - 1
		}
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("sqrt(%d) = %d\n", i, sqrt(i))
	}
}
