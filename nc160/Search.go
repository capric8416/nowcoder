package main

import "fmt"

/**
 *
 * @param A int整型一维数组
 * @param target int整型
 * @return int整型
 */
func search(nums []int, target int) int {
	// write code here

	for left, right := 0, len(nums)-1; left <= right; {
		center := left + (right-left)>>1
		if nums[center] == target {
			return center
		}

		if nums[left] <= nums[center] { // 左侧有序
			if nums[left] <= target && target < nums[center] { // 左区间
				right = center - 1
			} else { // 右区间
				left = center + 1
			}
		} else { // 右侧有序
			if nums[center] < target && target <= nums[right] {
				left = center + 1 // 右区间
			} else {
				right = center - 1 // 左区间
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{3, 1}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
	fmt.Println(search([]int{2, 4, 5, 6, 7, 0, 1}, 1))
	fmt.Println(search([]int{1, 2, 4, 5, 6, 7, 0}, 1))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, 1))
	fmt.Println(search([]int{7, 0, 1, 2, 4, 5, 6}, 1))
	fmt.Println(search([]int{6, 7, 0, 1, 2, 4, 5}, 1))
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 1))
}
