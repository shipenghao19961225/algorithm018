package leetcode

import "sort"

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	sort.Ints(nums)
	used := make([]int, len(nums))
	dfs(nums, used, []int{})
	return res
}

func dfs(nums, used, path []int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if (i != 0 && nums[i] == nums[i-1] && used[i-1] != 0) || (used[i] == 1) {
			continue
		}
		path = append(path, nums[i])
		used[i] = 1
		dfs(nums, used, path)
		path = path[:len(path)-1]
		used[i] = 0

	}
}
