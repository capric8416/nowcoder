package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return string字符串一维数组
 */
func Permutation(str string) []string {
	// write code here

	length := len(str)
	if length < 2 {
		return []string{str}
	}

	arr := []byte(str)

	result := make([]string, 0)
	for i, j := 0, 0; true; {
		result = append(result, string(arr))

		// 从排列的右端开始，找出第一个小于右边数字的序号i
		for i = length - 2; i >= 0 && arr[i] >= arr[i+1]; i-- {
		}

		if i < 0 {
			break
		}

		// 在a[i]的右侧区域中，找出所有大于a[i]的数字中最小的数字a[j]
		for j = length - 1; j > i && arr[j] <= arr[i]; j-- {
		}

		// 交换
		arr[i], arr[j] = arr[j], arr[i]

		// 反转, 得到下一个排列
		for left, right := i+1, length-1; left < right; left, right = left+1, right-1 {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	return result
}

func main() {
	fmt.Println(Permutation("abbc"))
}
