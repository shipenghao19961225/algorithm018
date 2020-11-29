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
