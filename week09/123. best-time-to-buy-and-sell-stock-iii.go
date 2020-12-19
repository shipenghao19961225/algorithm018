package leetcode

//T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
//T[i][k][1] = max(T[i-1][k][1], T[i-1][k-1][0] - prices[i]
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	s20, s21, s10, s11 := 0, -prices[0], 0, -prices[0]
	for i := 1; i < n; i++ {
		s20 = max(s20, s21+prices[i])
		s21 = max(s21, s10-prices[i])
		s10 = max(s10, s11+prices[i])
		s11 = max(s11, -prices[i])
	}
	return s20
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
