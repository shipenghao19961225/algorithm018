package leetcode

// 第一种为回溯，记忆化的方法
var cache []int

func climbStairs(n int) int {
	cache = make([]int, n+1)
	return dfs(0, n)
}

func dfs(cur, n int) int {
	if cur == n {
		return 1
	}
	if cache[cur] != 0 {
		return cache[cur]
	}
	result1 := dfs(cur+1, n)
	if cur+2 <= n {
		result1 += dfs(cur+2, n)
	}
	cache[cur] = result1
	return result1
}

// 第二种为常规动态规划
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	p, q := 1, 2
	for i := 3; i <= n; i++ {
		tmp := q
		q = p + q
		p = tmp
	}
	return q
}

// 变种
// 每次可以爬1到m级台阶

func climbStairs2(n, m int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= m && i-j >= 0; j++ {
			dp[i] += dp[i-j]
		}
	}
	return dp[n]
}

// 每次可以爬1到m级台阶，但是相邻之间的步数不能相同，这里还是需要通过填表的方法，解决动态规划问题
// 只不过步骤较多，较为繁琐

func climbStairs3(n, m int) int {
	if n <= 1 {
		return n
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m && i >= j; j++ {
			if i == j {
				dp[i][j] = 1
			}
			for k := 1; k <= m; k++ {
				if k == j {
					continue
				}
				dp[i][j] += dp[i-j][k]
			}
		}
	}
	res := 0
	for _, v := range dp[n] {
		res += v
	}
	return res
}
