package leetcode

import (
	"fmt"
	"math"
	"sort"
)

// 1. 最笨方法
// 	记忆化递归
var cache []int

func coinChange(coins []int, amount int) int {
	cache = make([]int, amount+1)
	return coinChange2(coins, amount)
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}
	if cache[amount] != 0 {
		return cache[amount]
	}
	minVal := math.MaxInt32
	for _, coin := range coins {
		result := coinChange2(coins, amount-coin)
		if result == -1 {
			continue
		}
		minVal = min(minVal, result)
	}
	if minVal == math.MaxInt32 {
		cache[amount] = -1
		return -1
	}
	cache[amount] = minVal + 1
	return minVal + 1
}

// 第二种方法
// 动态规划，参考国际站的代码，自己的代码写的不好看
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 第三种方法，疯狂的剪枝！！！！（类似贪心， 但不是严格意义的贪心， 但是剪枝很精妙）
var res int

func coinChange(coins []int, amount int) int {
	res = math.MaxInt32
	// 这里的sort有助于剪枝！
	sort.Slice(coins, func(a, b int) bool {
		return coins[a] > coins[b]
	})
	fmt.Println(coins)
	dfs(coins, amount, 0, 0)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
func dfs(coins []int, amount, count, cIndex int) {
	if amount == 0 {
		res = min(res, count)
		return
	}
	if cIndex >= len(coins) {
		return
	}
	// for循环中剪枝！
	for i := amount / coins[cIndex]; i >= 0 && i+count < res; i-- {
		dfs(coins, amount-i*coins[cIndex], count+i, cIndex+1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 4. 第四种解法
// 广度优先遍历
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	visited := make([]bool, amount+1)
	visited[amount] = true
	q := make([]int, 0, amount)
	q = append(q, amount)
	visited[amount] = true
	step := 1
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			front := q[0]
			q = q[1:]
			for _, coin := range coins {
				next := front - coin
				if next == 0 {
					return step
				} else if next < 0 {
					break
				} else {
					if !visited[next] {
						q = append(q, next)
						visited[next] = true
					}
				}
			}
		}
		step++
	}
	return -1
}
