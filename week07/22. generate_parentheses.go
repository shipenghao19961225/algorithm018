package leetcode

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	dfs("", 0, 0, 2*n)
	return res
}

func dfs(str string, left, right, depth int) {
	// termination
	if len(str) == depth {
		res = append(res, str)
		return
	}
	// process
	if left < depth/2 {
		// drill down
		dfs(str+"(", left+1, right, depth)
	}
	if right < left {
		// drill down
		dfs(str+")", left, right+1, depth)
	}
	// reverse state
}
