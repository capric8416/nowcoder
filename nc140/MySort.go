package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort(arr []int) []int {
	// write code here

	QuickSort(arr, 0, len(arr)-1)
	return arr
}

func InserationSort(arr []int, left int, right int) {
	for i := left; i <= right; i++ {
		// 以首个元素为基准
		pivot := arr[i]

		// 将左边比它大的都依次往右边移动
		var j int
		for j = i; j > left && arr[j-1] > pivot; j-- {
			arr[j] = arr[j-1]
		}

		arr[j] = pivot
	}
}

func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	// 少于10个元素使用插入排序
	if right-left < 9 {
		InserationSort(arr, left, right)
		return
	}

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

	pivot := arr[left] // 基准值

	// 从两端开始扫描，小的放左边(i)，大的放右边(j)，相等的放左(m)右(n)两端
	i, j, m, n := left, right, left, right
	for i < j {
		// j从右至左扫描，直到找到比基准值小的
		for ; j > i && arr[j] >= pivot; j-- {
			// 和基准相等的，交换到右边
			if arr[j] == pivot {
				arr[n], arr[j] = arr[j], arr[n]
				n-- // 位置被占了，左移
			}
		}
		// j位置元素放到左边i位置
		arr[i] = arr[j]

		// i从左至右扫描，直到找到比基准值大的
		for ; i < j && arr[i] <= pivot; i++ {
			// 和基准相等的，交换到左边
			if arr[i] == pivot {
				arr[m], arr[i] = arr[i], arr[m]
				m++ // 位置被占了，右移
			}
		}
		// i位置元素放到右边j位置
		arr[j] = arr[i]
	}
	// 空出来的i位置保存基准值
	arr[i] = pivot

	// 一趟结束之后，把和基准值一样的都移动到其周围
	// 左端
	for x, y := left, i-1; x < m; x, y = x+1, y-1 {
		arr[y], arr[x] = arr[x], arr[y]
	}
	// 右端
	for x, y := i+1, right; y > n; x, y = x+1, y-1 {
		arr[x], arr[y] = arr[y], arr[x]
	}

	// 处理左区间
	QuickSort(arr, left, i-1-(m-left))

	// 处理右区间
	QuickSort(arr, j+1+(right-n), right)
}

func CreateAscendingArray(count int) []int {
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = i
	}
	return slice
}

func CreateDescendingArray(count int) []int {
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = count - i - 1
	}
	return slice
}

func CreateRepeatingArray(count int) []int {
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = 0
	}
	return slice
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

func CompareArrays(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	times := 10000000

	fmt.Printf("===================== ramdom %v =====================\n", times)

	arr := CreateRandomArray(times)
	now := time.Now()
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("quick sort: ", time.Since(now))

	arr1 := CreateRandomArray(times)
	now = time.Now()
	sort.Ints(arr1)
	fmt.Println("sort ints: ", time.Since(now))

	fmt.Println("equals: ", CompareArrays(arr, arr1))

	fmt.Printf("===================== ascending %v =====================\n", times)

	arr = CreateAscendingArray(times)
	now = time.Now()
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("quick sort: ", time.Since(now))

	arr1 = CreateAscendingArray(times)
	now = time.Now()
	sort.Ints(arr1)
	fmt.Println("sort ints: ", time.Since(now))

	fmt.Println("equals: ", CompareArrays(arr, arr1))

	fmt.Printf("===================== descending %v =====================\n", times)

	arr = CreateDescendingArray(times)
	now = time.Now()
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("quick sort: ", time.Since(now))

	arr1 = CreateDescendingArray(times)
	now = time.Now()
	sort.Ints(arr1)
	fmt.Println("sort ints: ", time.Since(now))

	fmt.Println("equals: ", CompareArrays(arr, arr1))

	fmt.Printf("===================== repeating %v =====================\n", times)

	arr = CreateRepeatingArray(times)
	now = time.Now()
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("quick sort: ", time.Since(now))

	arr = CreateRepeatingArray(times)
	now = time.Now()
	sort.Ints(arr)
	fmt.Println("sort ints: ", time.Since(now))
}
