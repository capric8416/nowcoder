package main

import "fmt"

/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func solve(str string) string {
	// write code here

	length := len(str)

	str[0] = 'a'

	return str
}

func main() {
	fmt.Println(solve("012"))
	fmt.Println(solve("1234"))
	fmt.Println(solve("0123456789abcdefghijklmnopqrstuvwxyz"))
}
