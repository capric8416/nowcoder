package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func maxHeapSink(arr []int, parentIndex int) []int {
	length := len(arr)

	for childIndex := 2*parentIndex + 1; childIndex < length; childIndex = 2*childIndex + 1 {
		if childIndex+1 < length && arr[childIndex] < arr[childIndex+1] {
			childIndex++
		}

		if arr[parentIndex] >= arr[childIndex] {
			break
		}

		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]

		parentIndex = childIndex
	}

	return arr
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

	// 建立大根堆
	arr := input[:k]
	for i := (k - 1) / 2; i >= 0; i-- {
		arr = maxHeapSink(arr, i)
	}

	// 替换最大的元素并下沉
	for i := k; i < size; i++ {
		if input[i] < arr[0] {
			arr[0] = input[i]
			arr = maxHeapSink(arr, 0)
		}
	}

	return arr
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

	fmt.Println("MaxHeap: ", time.Since(now))
}
