package leetcode

// 1. 记忆化递归
var cache [][]int

func uniquePaths(m int, n int) int {
	cache = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		cache[i] = make([]int, n+1)
	}
	return dfs(1, 1, m, n)
}
func dfs(i, j, m, n int) int {
	if i == m && j == n {
		return 1
	}
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	result := 0
	if i < m {
		result += dfs(i+1, j, m, n)
	}
	if j < n {
		result += dfs(i, j+1, m, n)
	}
	cache[i][j] = result
	return result
}

// 2. dp
func uniquePaths(m int, n int) int {
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[m-1]
}
