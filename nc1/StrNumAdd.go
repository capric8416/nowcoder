package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	// write code here

	i, j := len(s)-1, len(t)-1

	var sum []byte
	if i > j {
		sum = make([]byte, i+2)
	} else {
		sum = make([]byte, j+2)
	}

	for x, k := byte(0), len(sum)-1; ; i, j, k = i-1, j-1, k-1 {
		if i >= 0 && j >= 0 {
			x = (s[i] - '0') + (t[j]) - '0' + sum[k]
		} else if i >= 0 {
			x = s[i] - '0' + sum[k]
		} else if j >= 0 {
			x = t[j] - '0' + sum[k]
		} else {
			break
		}

		sum[k-1], sum[k] = x/10, x%10+'0'
	}

	if sum[0] > 0 {
		sum[0] += '0'
		return string(sum[:])
	}

	return string(sum[1:])
}

func main() {
	fmt.Printf("99 + 9999 = %v\n", solve("99", "9999"))
	fmt.Printf("829 + 316 = %v\n", solve("829", "316"))
	fmt.Printf("9829 + 316 = %v\n", solve("9829", "316"))
	fmt.Printf("238942949829 + 316 = %v\n", solve("238942949829", "316"))
}
