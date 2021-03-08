package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func QuickSelect(arr []int, left int, right int, k int) []int {
	if left > right {
		return []int{}
	}
	if left == right {
		return arr[:left+1]
	}

	if diff := right - left; diff > 1 {
		// 三数取中
		center := left + (right-left)/2
		// 找出最大的值放在右边
		if arr[center] > arr[right] {
			arr[right], arr[center] = arr[center], arr[right]
		}
		if arr[left] > arr[right] {
			arr[right], arr[left] = arr[left], arr[right]
		}
		// 找出次大的值放在左边作为基准值
		if arr[center] > arr[left] {
			arr[center], arr[left] = arr[left], arr[center]
		}
	}

	pivot := arr[left] // 基准值

	// 从两端开始扫描，小的放左边(i)，大的放右边(j)
	i, j := left, right
	for i < j {
		// j从右至左扫描，直到找到比基准值小的
		for ; j > i && arr[j] >= pivot; j-- {
		}
		// j位置元素放到左边i位置
		arr[i] = arr[j]

		// i从左至右扫描，直到找到比基准值大的
		for ; i < j && arr[i] <= pivot; i++ {
		}
		// i位置元素放到右边j位置
		arr[j] = arr[i]
	}
	// 空出来的i位置保存基准值
	arr[i] = pivot

	if i > k {
		// 处理左区间
		return QuickSelect(arr, left, i-1, k)
	}
	if i < k {
		// 处理右区间
		return QuickSelect(arr, i+1, right, k)
	}

	return arr[:i+1]
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here

	size := len(input)
	if size == 0 || k == 0 || k > size {
		return []int{}
	}

	return QuickSelect(input, 0, size-1, k-1)
}

func CreateRandomArray(count int) []int {
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = i
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < count/2; i++ {
		j := r.Intn(count)
		if i != j {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	return slice
}

func main() {
	now := time.Now()

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 10; i < 10000; i++ {
		input := CreateRandomArray(i)

		k := r.Intn(6)
		for k == 0 {
			k = r.Intn(6)
		}

		result := GetLeastNumbers_Solution(input, k)
		if len(result) == 0 {
			fmt.Println("#109 fail")
			continue
		}

		sort.Ints(result)
		for index, value := range result {
			if index != value {
				fmt.Println("#116 fail")
				break
			}
		}
	}

	result := GetLeastNumbers_Solution([]int{4, 5, 1, 6, 2, 7, 3, 8, 4, 3, 2, 1, 1, 2, 3, 4}, 5)
	fmt.Println(result)

	fmt.Println("QuickSelect: ", time.Since(now))
}
