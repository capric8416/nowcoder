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

	sLength, tLength := len(s), len(t)
	mulLength := sLength + tLength
	mul := make([]byte, mulLength)

	var v byte
	for i, offset := tLength-1, 0; i >= 0; i, offset = i-1, offset+1 {
		for j, k := sLength-1, mulLength-1-offset; j >= 0; j, k = j-1, k-1 {
			v = (t[i]-'0')*(s[j]-'0') + mul[k]
			mul[k] = v % 10
			if v >= 10 {
				mul[k-1] += v / 10
				if mul[k-1] >= 10 {
					mul[k-2] += mul[k-1] / 10
					mul[k-1] %= 10
				}
			}
		}
	}

	if mul[0] == 0 {
		mulLength--
		mul = mul[1:]
	}
	for i := 0; i < mulLength; i++ {
		mul[i] += '0'
	}

	return string(mul)
}

func main() {
	fmt.Println(solve("1", "9"))
	fmt.Println(solve("1", "99"))
	fmt.Println(solve("99", "9999"))
	fmt.Println(solve("12345", "999999"))
}
