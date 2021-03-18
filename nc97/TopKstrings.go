package main

import (
	"fmt"
	"strconv"
)

type KeyFrequency struct {
	Key       string
	Frequency int
}

func (kv1 KeyFrequency) lt(kv2 KeyFrequency) bool {
	if kv1.Frequency < kv2.Frequency {
		return true
	}
	if kv1.Frequency == kv2.Frequency && kv1.Key > kv2.Key {
		return true
	}
	return false
}

func (kv1 KeyFrequency) gt(kv2 KeyFrequency) bool {
	if kv1.Frequency > kv2.Frequency {
		return true
	}
	if kv1.Frequency == kv2.Frequency && kv1.Key < kv2.Key {
		return true
	}
	return false
}

func maxHeapSink(arr []KeyFrequency, parentIndex int, length int) {
	for childIndex := 2*parentIndex + 1; childIndex < length; childIndex = 2*childIndex + 1 {
		if childIndex+1 < length && arr[childIndex].lt(arr[childIndex+1]) {
			childIndex++
		}

		if arr[parentIndex].gt(arr[childIndex]) {
			break
		}

		arr[parentIndex], arr[childIndex] = arr[childIndex], arr[parentIndex]

		parentIndex = childIndex
	}
}

/**
 * return topK string
 * @param strings string字符串一维数组 strings
 * @param k int整型 the k
 * @return string字符串二维数组
 */
func topKstrings(strings []string, k int) [][]string {
	// write code here

	// 统计频率
	strFrequency := make(map[string]int)
	for _, k := range strings {
		strFrequency[k]++
	}

	// 建立初始堆
	length := 0
	heap := make([]KeyFrequency, len(strFrequency))
	for k, v := range strFrequency {
		heap[length] = KeyFrequency{Key: k, Frequency: v}
		length++
	}

	// 调整为大根堆
	for i := (length - 1) / 2; i >= 0; i-- {
		maxHeapSink(heap, i, length)
	}

	result := make([][]string, k)
	for i, last := 0, length-1; i < k; i, last = i+1, last-1 {
		result[i] = []string{heap[0].Key, strconv.Itoa(heap[0].Frequency)}

		// 调整为大根堆
		heap[0], heap[last] = heap[last], heap[0]
		maxHeapSink(heap, 0, last)
	}

	return result
}

func main() {
	fmt.Println(topKstrings([]string{"1", "2", "3", "4"}, 2))
	fmt.Println(topKstrings([]string{"1", "1", "2", "3"}, 2))
	fmt.Println(topKstrings([]string{"1", "1", "2", "3", "5", "8", "0", "1", "1", "2", "3", "5", "8", "9", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, 5))
}
