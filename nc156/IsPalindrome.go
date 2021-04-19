package main

import (
	"fmt"
)

/**
 *
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	// write code here

	if x < 0 {
		return false
	}

	var first, mul int
	for x > 0 {
		if x < 10 {
			return true
		}

		for first, mul = x, 1; first > 10; {
			first /= 10
			mul *= 10
		}
		last := x % 10
		if first != last {
			return false
		}

		x -= first * mul
		x /= 10
	}

	return true
}

// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}

// 	var t, digits int
// 	for t, digits = x, 0; t != 0; t = t / 10 {
// 		digits++
// 	}

// 	count := digits / 2
// 	var j, low, high int
// 	for i := 0; i < count; i++ {
// 		for low, j = 1, i; j > 0; j-- {
// 			low *= 10
// 		}
// 		low = x / low % 10
// 		for high, j = 1, digits-i-1; j > 0; j-- {
// 			high *= 10
// 		}
// 		high = x / high % 10
// 		if low != high {
// 			return false
// 		}
// 	}

// 	return true
// }

func main() {
	fmt.Println(isPalindrome(9223372036854775807))
	fmt.Println(isPalindrome(922337203685477580))
	fmt.Println(isPalindrome(92233720368547758))
	fmt.Println(isPalindrome(9223372036854775))
	fmt.Println(isPalindrome(922337203685477))
	fmt.Println(isPalindrome(92233720368547))
	fmt.Println(isPalindrome(9223372036854))
	fmt.Println(isPalindrome(922337203685))
	fmt.Println(isPalindrome(92233720368))
	fmt.Println(isPalindrome(9223372036))
	fmt.Println(isPalindrome(922337203))
	fmt.Println(isPalindrome(92233720))
	fmt.Println(isPalindrome(9223372))
	fmt.Println(isPalindrome(922337))
	fmt.Println(isPalindrome(92233))
	fmt.Println(isPalindrome(9223))
	fmt.Println(isPalindrome(922))
	fmt.Println(isPalindrome(92))
	fmt.Println(isPalindrome(9))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(12344321))
	fmt.Println(isPalindrome(1234554321))
	fmt.Println(isPalindrome(123456654321))
	fmt.Println(isPalindrome(12345677654321))
	fmt.Println(isPalindrome(1234567887654321))
	fmt.Println(isPalindrome(123456789987654321))
	fmt.Println(isPalindrome(1234567890987654321))
	fmt.Println(isPalindrome(1234))
	fmt.Println(isPalindrome(565))
	fmt.Println(isPalindrome(123421))
}
