package main

import "fmt"

/**
 * 进制转换
 * @param M int整型 给定整数
 * @param N int整型 转换到的进制
 * @return string字符串
 */
func solve(M int, N int) string {
	// write code here

	sign := false
	if M < 0 {
		M = -M
		sign = true
	}

	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	var result string
	for ; M/N != 0; M /= N {
		result = hex[M%N] + result
	}
	result = hex[M%N] + result
	if sign {
		result = "-" + result
	}

	return result
}

func main() {
	fmt.Println(solve(156, 2))
	fmt.Println(solve(255, 16))
	fmt.Println(solve(-156, 2))
	fmt.Println(solve(-255, 16))
}
