package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	preProfit0 := 0
	profit0, profit1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		nextProfit0 := max(profit0, profit1+prices[i])
		nextProfit1 := max(profit1, preProfit0-prices[i])
		preProfit0 = profit0
		profit0, profit1 = nextProfit0, nextProfit1
	}
	return profit0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])
// T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k][0] - prices[i])

// T[i][0] = max(T[i - 1][0], T[i - 1][1] + prices[i])
// T[i][1] = max(T[i - 1][1], T[i - 2][0] - prices[i])
