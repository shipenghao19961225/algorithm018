package leetcode

// 动态规划
func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = i
	}
	for i := 1; i <= len(word2); i++ {
		temp := dp[0]
		dp[0] = i
		for j := 1; j <= len(word1); j++ {
			cur := temp
			temp = dp[j]
			if word1[j-1] == word2[i-1] {
				dp[j] = cur
			} else {
				dp[j] = min(cur, dp[j], dp[j-1]) + 1
			}

		}

	}
	return dp[len(word1)]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

// 记忆化递归，时间复杂度要小于dp，精妙
var cache [][]int

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	cache = make([][]int, m+1)
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, n+1)
	}
	return dfs(m, n, word1, word2)
}

func dfs(p1, p2 int, word1, word2 string) int {
	if cache[p1][p2] > 0 {
		return cache[p1][p2]
	}
	if p1*p2 == 0 {
		return p1 + p2
	}
	if word1[p1-1] == word2[p2-1] {
		cache[p1][p2] = dfs(p1-1, p2-1, word1, word2)
	} else {
		cache[p1][p2] = 1 + min(dfs(p1-1, p2-1, word1, word2), dfs(p1-1, p2, word1, word2), dfs(p1, p2-1, word1, word2))

	}
	return cache[p1][p2]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else {
		if b < c {
			return b
		}
		return c
	}
}
