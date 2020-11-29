package leetcode

// https://leetcode-cn.com/problems/house-robber/solution/da-jia-jie-she-dong-tai-gui-hua-jie-gou-hua-si-lu-/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre, cur := 0, 0
	for i := 0; i < len(nums); i++ {
		tmp := cur
		cur = max(pre+nums[i], cur)
		pre = tmp
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
