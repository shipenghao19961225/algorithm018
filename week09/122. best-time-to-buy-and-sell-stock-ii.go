package leetcode

// 贪心解法
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// dp解法
//T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
//T[i][k][1] = max(T[i-1][k][1], T[i-1][k-1][0] - prices[i]
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	s0, s1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		newS0 := max(s0, s1+prices[i])
		newS1 := max(s1, s0-prices[i])
		s0, s1 = newS0, newS1
	}
	return s0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
