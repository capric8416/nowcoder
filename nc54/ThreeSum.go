package main

import "fmt"

func InserationSort(nums []int, left int, right int) {
	for i := left; i <= right; i++ {
		// 以首个元素为基准
		pivot := nums[i]

		// 将左边比它大的都依次往右边移动
		var j int
		for j = i; j > left && nums[j-1] > pivot; j-- {
			nums[j] = nums[j-1]
		}

		nums[j] = pivot
	}
}

func QuickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	// 少于10个元素使用插入排序
	if right-left < 9 {
		InserationSort(nums, left, right)
		return
	}

	// 三数取中
	center := left + (right-left)/2
	// 找出最大的值放在右边
	if nums[center] > nums[right] {
		nums[right], nums[center] = nums[center], nums[right]
	}
	if nums[left] > nums[right] {
		nums[right], nums[left] = nums[left], nums[right]
	}
	// 找出次大的值放在左边作为基准值
	if nums[center] > nums[left] {
		nums[center], nums[left] = nums[left], nums[center]
	}

	pivot := nums[left] // 基准值

	// 从两端开始扫描，小的放左边(i)，大的放右边(j)，相等的放左(m)右(n)两端
	i, j, m, n := left, right, left, right
	for i < j {
		// j从右至左扫描，直到找到比基准值小的
		for ; j > i && nums[j] >= pivot; j-- {
			// 和基准相等的，交换到右边
			if nums[j] == pivot {
				nums[n], nums[j] = nums[j], nums[n]
				n-- // 位置被占了，左移
			}
		}
		// j位置元素放到左边i位置
		nums[i] = nums[j]

		// i从左至右扫描，直到找到比基准值大的
		for ; i < j && nums[i] <= pivot; i++ {
			// 和基准相等的，交换到左边
			if nums[i] == pivot {
				nums[m], nums[i] = nums[i], nums[m]
				m++ // 位置被占了，右移
			}
		}
		// i位置元素放到右边j位置
		nums[j] = nums[i]
	}
	// 空出来的i位置保存基准值
	nums[i] = pivot

	// 一趟结束之后，把和基准值一样的都移动到其周围
	// 左端
	for x, y := left, i-1; x < m; x, y = x+1, y-1 {
		nums[y], nums[x] = nums[x], nums[y]
	}
	// 右端
	for x, y := i+1, right; y > n; x, y = x+1, y-1 {
		nums[x], nums[y] = nums[y], nums[x]
	}

	// 处理左区间
	QuickSort(nums, left, i-1-(m-left))

	// 处理右区间
	QuickSort(nums, j+1+(right-n), right)
}

func BinarySearchAny(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		center := left + (right-left)/2
		if nums[center] < target {
			left = center + 1
		} else if nums[center] > target {
			right = center - 1
		} else {
			return center
		}
	}

	return left
}

func BinarySearchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		center := left + (right-left)/2
		if nums[center] < target {
			left = center + 1
		} else if nums[center] > target {
			right = center - 1
		} else {
			right = center
		}
	}

	return left
}

func BinarySearchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		center := left + (right-left)/2
		if nums[center] < target {
			left = center + 1
		} else if nums[center] > target {
			right = center - 1
		} else {
			left = center
			if right-left == 1 {
				return left
			}
		}
	}

	return left
}

/**
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func threeSum(num []int) [][]int {
	// write code here

	right := len(num) - 1

	QuickSort(num, 0, right)

	minZeroPos, maxZeroPos, minNegativePos, maxNegativePos, minPositivePos, maxPositivePos := -1, -1, -1, -1, -1, -1

	if num[0] < 0 {
		minNegativePos = 0
	}

	if num[right] > 0 {
		maxPositivePos = right
	}

	n := BinarySearchAny(num, 0)
	if num[n] == 0 {
		minZeroPos = BinarySearchLeft(num[:n+1], 0)
		maxNegativePos = minZeroPos - 1

		maxZeroPos = BinarySearchRight(num[minZeroPos:], 0)
		maxZeroPos += n
		if maxZeroPos < right {
			minPositivePos = maxZeroPos + 1
		}
	} else {
		if num[n] < 0 {
			maxNegativePos = n
			minPositivePos = n + 1
		} else {
			minPositivePos = n
			maxNegativePos = n - 1
		}
	}

	result := make([][]int, 0)

	// 全为零
	if maxZeroPos-minZeroPos > 1 {
		result = append(result, []int{0, 0, 0})
	}

	if maxNegativePos == -1 || minPositivePos == -1 {
		return result
	}

	// 一个正数，一个负数
	if n != -1 {
		for i := minNegativePos; i <= maxNegativePos; {
			target := -num[i]
			j := BinarySearchAny(num[minPositivePos:maxPositivePos+1], target)
			if num[j] == target {
				result = append(result, []int{num[i], 0, num[minPositivePos+j]})
			}

			current := num[i]
			for i <= right && num[i] == current {
				i++
			}
		}
	}

	// 两个正数，一个负数
	for i := minPositivePos; i < maxPositivePos; {
		for j := i + 1; j <= maxPositivePos; {
			target := -(num[i] + num[j])
			k := minNegativePos + BinarySearchAny(num[minNegativePos:maxNegativePos+1], target)
			if num[k] == target {
				result = append(result, []int{num[k], num[i], num[j]})
			}

			current := num[j]
			for j <= right && num[j] == current {
				j++
			}
		}

		current := num[i]
		for i <= right && num[i] == current {
			i++
		}
	}

	// 选两个负数，一个正数
	for i := minNegativePos; i < maxNegativePos; {
		for j := i + 1; j <= maxNegativePos; {
			target := -(num[i] + num[j])
			k := minPositivePos + BinarySearchAny(num[minPositivePos:maxPositivePos+1], target)
			if num[k] == target {
				result = append(result, []int{num[i], num[j], num[k]})
			}

			current := num[j]
			for j <= maxPositivePos && num[j] == current {
				j++
			}
		}

		current := num[i]
		for i <= maxPositivePos && num[i] == current {
			i++
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-10, 0, 10, 20, -10, -40}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum([]int{-2, -1, -1, -1, -3, -4, -5, -1, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 4, 5}))
	fmt.Println(threeSum([]int{-2, -1, -5, -1, -1, 0, 0, 0, 0, 1, 1, 1, 4, 5}))
	fmt.Println(threeSum([]int{-2, -1, -1, -1, -3, -4, -5, -1, -1, 1, 1, 1, 1, 1, 2, 3, 3, 4, 4, 5}))
	fmt.Println(threeSum([]int{-2, -1, -1, -1, -3, -4, -5, -1, 1, 1, 1, 1, 1, 2, 3, 3, 4, 4, 5}))
}
