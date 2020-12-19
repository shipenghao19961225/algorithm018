学习笔记

## 爬楼梯问题总结

|                         | 1     | 2     | 3    | 4    | 5    |
| ----------------------- | ----- | ----- | ---- | ---- | ---- |
| 70. 爬楼梯（题目变形）  | 12.15 |       |      |      |      |
| 746. 使用最小花费爬楼梯 | 12.15 | 12.15 |      |      |      |
| 121. 股票1              | 12.16 | 12.16 |      |      |      |
| 122. 股票2              | 12.16 | 12.16 |      |      |      |
| 123.股票3               | 12.16 | 12.16 |      |      |      |
| 124.股票4               | 12.16 | 12.16 |      |      |      |
| 309.股票带冷却          | 12.16 | 12.16 |      |      |      |
| 714.股票带手续费        | 12.16 | 12.16 |      |      |      |
| 387. 字符串唯一字符     | 12.19 | 12.19 |      |      |      |
| 300. 最长上升子序列     | 12.19 | 12.19 |      |      |      |
| 91. 解码方法            | 12.19 | 12.19 |      |      |      |
| 32. 最长有效括号        | 12.19 | 12.19 |      |      |      |

```go
package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(4))
	fmt.Println(climbStairs2(4, 4))
	fmt.Println(climbStairs3(5,5))
}

// 爬楼汇总

// 每次爬一级台阶，爬两级台阶
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

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
	dp := make([][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m && i >= j; j++ {
			if i == j {
				dp[i][j] = 1
			}
			for k := 1; k <= m; k++{
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

```





这周作业没做完。