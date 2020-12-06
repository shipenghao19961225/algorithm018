package leetcode

import "strings"

var res [][]string
var col map[int]bool
var pie map[int]bool
var na map[int]bool

func solveNQueens(n int) [][]string {
	res = [][]string{}
	col = make(map[int]bool)
	pie = make(map[int]bool)
	na = make(map[int]bool)
	dfs(0, n, []int{})
	return res
}

func dfs(row, n int, path []int) {
	// termination
	if row == n {
		res = append(res, generate(path, n))
		return
	}

	for i := 0; i < n; i++ {
		// process
		if col[i] || pie[row-i] || na[row+i] {
			continue
		}
		col[i] = true
		pie[row-i] = true // 这里的对角线非常巧妙，行加列或者行减列就能标识对角线
		na[row+i] = true  //
		path = append(path, i)
		// drill down
		dfs(row+1, n, path)
		col[i] = false
		pie[row-i] = false
		na[row+i] = false
		// reverse
		path = path[:len(path)-1]
	}
}
func generate(path []int, n int) []string {
	ans := []string{}
	for _, v := range path {
		tmp := strings.Repeat(".", v) + "Q" + strings.Repeat(".", n-v-1)
		ans = append(ans, tmp)
	}
	return ans
}
