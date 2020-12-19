package leetcode

// 题解见，讲的透彻明了
// https://leetcode-cn.com/circle/article/qiAgHn/
//T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
//T[i][k][1] = max(T[i-1][k][1], T[i-1][k-1][0] - prices[i])
//
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	s0, s1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		s0 = max(s0, s1+prices[i])
		s1 = max(s1, -prices[i])
	}
	return s0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
