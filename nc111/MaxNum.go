package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
 * 最大数
 * @param nums int整型一维数组
 * @return string字符串
 */
func solve(nums []int) string {
	// write code here

	length := len(nums)
	n := make([]string, len(nums))
	for i := 0; i < length; i++ {
		n[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(n, func(i, j int) bool {
		return n[i]+n[j] > n[j]+n[i]
	})

	if n[0] == "0" {
		return "0"
	}
	return strings.Join(n, "")
}

func main() {
	fmt.Println(solve([]int{3, 30, 34, 5, 9}))
}
