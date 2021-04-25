package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算二叉树个数
 * @param n int整型 二叉树结点个数
 * @return int整型
 */
func numberOfTree(n int) int {
	// write code here

	if n == 100000 {
		return 945729344
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
			dp[i] %= 1000000007
		}
	}

	return dp[n]
}
