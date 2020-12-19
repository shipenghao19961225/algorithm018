package leetcode

// 这道题三种方法

// 1. 栈，方法，始终维护栈的最底端是最后一个未匹配括号的索引

func longestValidParentheses(s string) int {
	stack := make([]int, 0, len(s)+1)
	stack = append(stack, -1)
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
				continue
			}
			res = max(res, i-stack[len(stack)-1])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2， 动态规划，难点在于如何写出正确的动态转移方程，当然，这里的加一个空字符串的做法值得学习
func longestValidParentheses(s string) int {
	s = " " + s
	dp := make([]int, len(s))
	res := 0
	for i := 2; i < len(dp); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			} else if s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//  3. 脑筋急转弯做法
func longestValidParentheses(s string) int {
	left, right := 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if right == left {
			res = max(res, 2*left)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left == right {
			res = max(res, 2*right)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
