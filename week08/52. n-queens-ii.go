package leetcode

func totalNQueens(n int) int {
	res := 0
	dfs(0, n, 0, 0, 0, &res)
	return res
}
func dfs(row, n, cols, pies, nas int, res *int) {
	if row == n {
		*res++
	}
	avaPos := (1<<n - 1) & (^(cols | pies | nas))
	for avaPos > 0 {
		Pos := avaPos & (-avaPos)
		avaPos &= avaPos - 1
		dfs(row+1, n, cols|Pos, (pies|Pos)<<1, (nas|Pos)>>1, res)
	}
}
