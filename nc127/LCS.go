package main

import "fmt"

/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS(str1 string, str2 string) string {
	// write code here

	m, n := len(str1), len(str2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	str2Index := 0
	maxCommonLength := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if maxCommonLength < dp[i+1][j+1] {
					str2Index = j + 1
					maxCommonLength = dp[i+1][j+1]
				}
			}
		}
	}

	if maxCommonLength == 0 {
		return ""
	}
	return str2[str2Index-maxCommonLength : str2Index]
}

func main() {
	fmt.Println(LCS("2LQ74WK8Ld0x7d8FP8l61pD7Wsz1E9xOMp920hM948eGjL9Kb5KJt80", "U08U29zzuodz16CBZ8xfpmmn5SKD80smJbK83F2T37JRqYfE76vh6hrE451uFQ100ye9hog1Y52LDk0L52SuD948eGjLz0htzd5YF9J1Y6oI7562z4T2"))
}
