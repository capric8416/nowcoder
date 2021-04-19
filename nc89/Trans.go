package main

import (
	"fmt"
)

/**
 *
 * @param s string字符串
 * @param n int整型
 * @return string字符串
 */
func trans(s string, n int) string {
	// write code here

	result := make([]byte, n)

	for i, pos := n-1, 0; i >= 0; {
		temp := make([]byte, 0)

		var j int
		for j = i; j >= 0 && s[j] != ' '; j-- {
			if 'A' <= s[j] && s[j] <= 'Z' {
				temp = append(temp, s[j]+32)
			} else {
				temp = append(temp, s[j]-32)
			}
		}

		count := len(temp)
		for k := count - 1; k >= 0; k-- {
			result[pos] = temp[k]
			pos++
		}

		if pos < n {
			result[pos] = ' '
			pos++
		}

		i = j - 1
	}

	return string(result)
}

func main() {
	fmt.Println(trans("This is a sample", 16))
}
