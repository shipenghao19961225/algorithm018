package leetcode

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if k < 0 || n < k {
		return res
	}
	dfs(n, k, 1, []int{})
	return res
}

func dfs(n, k, begin int, path []int) {
	// termination
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	// process
	// 这里的剪枝很精髓！！！！学习一下
	// 因为当前路径一定有一定长度了，后面需要补全 k - len(path)个元素的就可以了
	// 所以索引到i <= n-(k-len(path))+1就可以，再往后元素就不够用，浪费时间。
	for i := begin; i <= n-(k-len(path))+1; i++ {

		path = append(path, i)
		// drill down
		dfs(n, k, i+1, path)
		// reverse the state if needed
		path = path[:len(path)-1]
	}
}
