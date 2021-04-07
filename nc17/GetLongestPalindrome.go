package main

import "fmt"

func centerToBothSides(A string, n int, left int, right int) int {
	for left >= 0 && right < n && A[left] == A[right] {
		left--
		right++
	}
	return right - left - 1
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param A string字符串
 * @param n int整型
 * @return int整型
 */
func getLongestPalindrome(A string, n int) int {
	// write code here

	length := 0
	for i := 0; i < n; i++ {
		count := centerToBothSides(A, n, i, i)
		if count > length {
			length = count
		}
		count = centerToBothSides(A, n, i, i+1)
		if count > length {
			length = count
		}
	}

	return length
}

func main() {
	fmt.Println(getLongestPalindrome("cbc", 3))
	fmt.Println(getLongestPalindrome("baabccc", 7))
}
