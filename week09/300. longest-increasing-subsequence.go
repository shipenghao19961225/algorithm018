package leetcode

func lengthOfLIS(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
