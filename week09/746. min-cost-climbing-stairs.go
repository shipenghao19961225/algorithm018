package leetcode

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	n := len(cost)
	pre, cur := cost[0], cost[1]
	for i := 2; i < n; i++ {
		cur, pre = min(pre, cur)+cost[i], cur
	}
	return min(cur, pre)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
