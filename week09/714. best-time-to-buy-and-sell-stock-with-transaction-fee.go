package leetcode

func maxProfit(prices []int, fee int) int {
	// T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i] - fee)
	// T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k][0] - prices[i])
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	s0, s1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		newS0 := max(s0, s1+prices[i]-fee)
		newS1 := max(s1, s0-prices[i])
		s0 = newS0
		s1 = newS1
	}
	return s0

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
