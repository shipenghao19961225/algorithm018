package leetcode

// 从前往后进行贪心
func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums) && i <= maxReach; i++ {
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 从后往前进行贪心， 不能提前终止，貌似代码更简洁
func canJump(nums []int) bool {
	last := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last == 0
}
