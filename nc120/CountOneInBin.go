package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param n int整型
 * @return int整型
 */
func NumberOf1(n int) int {
	// write code here

	// cnt := 0
	// for n != 0 {
	//     n = n & int(int32(n)-1)
	//     cnt ++
	// }

	// return cnt

	// count := 0
	// if n >= 0 {
	//     for n != 0 {
	//         count++
	//         n = n & (n -1)
	//     }
	// } else {
	//     n = n & 0xffffffff
	//     for n != 0 {
	//         count++
	//         n = n & (n -1)
	//     }
	// }
	// return count

	count := 32

	if n >= 0 {
		ones := 0
		for true {
			if n%2 == 1 {
				ones++
			}

			n /= 2
			if n == 0 {
				break
			}
		}
		return ones
	} else {
		n = -n
		bits := make([]byte, 0)
		mod, ones := 0, 0
		for true {
			mod = n % 2
			bits = append(bits, byte(mod))
			if mod == 1 {
				ones++
			}

			n /= 2
			if n == 0 {
				break
			}
		}

		if bits[0] == 1 {
			ones = count - ones + 1
		} else {
			ones = 0
			add := true
			for i := len(bits); i < count; i++ {
				bits = append(bits, 0)
			}
			for i := 0; i < 32; i++ {
				bits[i] = 1 - bits[i]
				if add {
					if bits[i] == 1 {
						bits[i] = 0
					} else {
						bits[i] = 1
						add = false
					}
				}
				ones += int(bits[i])
			}
		}

		return ones
	}
}

func main() {
	fmt.Println(NumberOf1(10))
	fmt.Println(NumberOf1(0))
	fmt.Println(NumberOf1(1))
	fmt.Println(NumberOf1(-2147483648))
	fmt.Println(NumberOf1(-1))
}
