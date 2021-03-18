package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
func judge(str string) bool {
	// write code here

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(judge("121"))
	fmt.Println(judge("1211"))
	fmt.Println(judge("123454321"))
	fmt.Println(judge("12343321"))
	fmt.Println(judge("1234554321"))
	fmt.Println(judge("1234564321"))
}
