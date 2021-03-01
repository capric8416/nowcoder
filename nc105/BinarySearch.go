package main

import "fmt"

/**
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
 */
func upper_bound_(n int, v int, a []int) int {
	// write code here

	// 最大值任然小于目标值
	if a[n-1] < v {
		return n + 1
	}

	left, right := 0, n-1
	for left < right {
		center := left + (right-left)/2
		if a[center] < v {
			// 如果中值小于目标，则center左侧区间不可能是目标，裁剪掉
			left = center + 1
		} else {
			// 如果中值大于等于目标，则center右侧区间不可能是目标，裁剪掉，但自身可能是目标
			right = center
		}
	}
	return left + 1
}

func main() {
	fmt.Println("========================================")
	arr := []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5}
	pos := upper_bound_(len(arr), 4, arr)
	fmt.Println(pos, "->", arr[pos])

	fmt.Println("========================================")
	arr = []int{1, 2, 3, 4, 4, 5}
	pos = upper_bound_(len(arr), 4, arr)
	fmt.Println(pos, "->", arr[pos])

	fmt.Println("========================================")
	arr = []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 6}
	pos = upper_bound_(len(arr), 4, arr)
	fmt.Println(pos, "->", arr[pos])

	fmt.Println("========================================")
	arr = []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 6}
	pos = upper_bound_(len(arr), 7, arr)
	fmt.Println(pos)

	fmt.Println("========================================")
	arr = []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	pos = upper_bound_(len(arr), 4, arr)
	fmt.Println(pos)
}
