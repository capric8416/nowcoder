package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func InserationSort(arr []*Interval, left int, right int) {
	for i := left; i <= right; i++ {
		// 以首个元素为基准
		pivot := arr[i]

		// 将左边比它大的都依次往右边移动
		var j int
		for j = i; j > left && arr[j-1].Start > pivot.Start; j-- {
			arr[j] = arr[j-1]
		}

		arr[j] = pivot
	}
}

func QuickSort(arr []*Interval, left int, right int) {
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
	if arr[center].Start > arr[right].Start {
		arr[right], arr[center] = arr[center], arr[right]
	}
	if arr[left].Start > arr[right].Start {
		arr[right], arr[left] = arr[left], arr[right]
	}
	// 找出次大的值放在左边作为基准值
	if arr[center].Start > arr[left].Start {
		arr[center], arr[left] = arr[left], arr[center]
	}

	pivot := arr[left] // 基准值

	// 从两端开始扫描，小的放左边(i)，大的放右边(j)，相等的放左(m)右(n)两端
	i, j, m, n := left, right, left, right
	for i < j {
		// j从右至左扫描，直到找到比基准值小的
		for ; j > i && arr[j].Start >= pivot.Start; j-- {
			// 和基准相等的，交换到右边
			if arr[j] == pivot {
				arr[n], arr[j] = arr[j], arr[n]
				n-- // 位置被占了，左移
			}
		}
		// j位置元素放到左边i位置
		arr[i] = arr[j]

		// i从左至右扫描，直到找到比基准值大的
		for ; i < j && arr[i].Start <= pivot.Start; i++ {
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

/**
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func merge(intervals []*Interval) []*Interval {
	// write code here

	length := len(intervals)
	if length < 2 {
		return intervals
	}

	QuickSort(intervals, 0, length-1)

	results := []*Interval{intervals[0]}
	for i := 1; i < length; i++ {
		curr := intervals[i]
		prev := results[len(results)-1]
		// 不包含
		if prev.End < curr.Start {
			results = append(results, curr)
		}
		// 全包含
		if prev.End >= curr.End {
			continue
		}
		// 部分包含
		if prev.End >= curr.Start && prev.End < curr.End {
			prev.End = curr.End
			continue
		}
	}

	return results
}

func main() {
	intervals := []*Interval{&Interval{Start: 20, End: 60}, &Interval{Start: 10, End: 30}, &Interval{Start: 80, End: 100}, &Interval{Start: 150, End: 180}}
	fmt.Println(merge(intervals))
}
