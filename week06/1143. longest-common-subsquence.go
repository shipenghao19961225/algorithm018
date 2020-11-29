package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text1)+1)
	m := len(text2) + 1
	n := len(text1) + 1

	for i := 1; i < m; i++ {
		follow := dp[0]
		for j := 1; j < n; j++ {
			tmp := follow
			follow = dp[j]
			if text1[j-1] == text2[i-1] {
				dp[j] = tmp + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
