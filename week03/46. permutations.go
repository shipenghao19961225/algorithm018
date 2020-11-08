package leetcode

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	used := make([]int, len(nums))
	dfs(nums, []int{}, used)
	return res
}

func dfs(nums, path, used []int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == 0 {
			path = append(path, nums[i])
			used[i] = 1
			dfs(nums, path, used)
			path = path[:len(path)-1]
			used[i] = 0
		}

	}
}
