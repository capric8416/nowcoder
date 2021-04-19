package main

import "fmt"

/**
 *
 * @param x int整型
 * @return int整型
 */
func reverse(x int) int {
	// write code here

	if -10 < x && x < 10 {
		return x
	}

	negative := false
	if x < 0 {
		negative = true
	}

	result := 0
	for x != 0 {
		digit := int32(x % 10)

		newResult := int32(result)*10 + digit
		if (!negative && newResult < 0) || (negative && newResult > 0) || (newResult-digit)/10 != int32(result) {
			return 0
		}
		result = int(newResult)

		x /= 10
	}

	return result
}

func main() {
	fmt.Println(reverse(0))
	fmt.Println(reverse(9))
	fmt.Println(reverse(-9))
	fmt.Println(reverse(10))
	fmt.Println(reverse(123456789000000000))
	fmt.Println(reverse(-123456789000000000))
	fmt.Println(reverse(7463847412))
	fmt.Println(reverse(-8463847412))
	fmt.Println(reverse(11123456789))
	fmt.Println(reverse(1123456789))
	fmt.Println(reverse(7463847413))
	fmt.Println(reverse(7863847412))
	fmt.Println(reverse(8463847412))
	fmt.Println(reverse(-8463847413))
	fmt.Println(reverse(-8465847412))
	fmt.Println(reverse(-9463847412))
}
