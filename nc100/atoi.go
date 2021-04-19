package main

import "fmt"

/**
 *
 * @param str string字符串
 * @return int整型
 */
func atoi(str string) int {
	// write code here

	length := len(str)
	if length == 0 {
		return 0
	}

	var space int
	for space = 0; space < length; {
		if str[space] == ' ' {
			space++
		} else {
			break
		}
	}
	if space > 0 {
		str = str[space:]
		length -= space
	}

	negative := false
	if str[0] == '+' {
		length--
		str = str[1:]
	} else if str[0] == '-' {
		length--
		str = str[1:]
		negative = true
	}

	result := 0
	for i := 0; i < length; i++ {
		if str[i] < '0' || str[i] > '9' {
			return result
		}

		digit := int32(str[i] - '0')
		if negative {
			digit = -digit
		}

		newResult := int32(result)*10 + digit
		if (!negative && newResult < 0) || (negative && newResult > 0) || (newResult-digit)/10 != int32(result) {
			if negative {
				return -2147483648
			}
			return 2147483647
		}
		result = int(newResult)
	}

	return result
}

func main() {
	fmt.Println(atoi("123"))
	fmt.Println(atoi("-123"))
	fmt.Println(atoi("      -117e40091539"))
	fmt.Println(atoi("    11333713950"))
}
