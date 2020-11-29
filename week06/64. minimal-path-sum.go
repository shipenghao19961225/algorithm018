package leetcode

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = grid[0][i] + dp[i-1]
	}
	for i := 1; i < m; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j]+grid[i][j], dp[j-1]+grid[i][j])
		}
	}
	return dp[n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
