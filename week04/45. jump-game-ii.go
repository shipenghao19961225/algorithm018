package leetcode

// 在遍历数组时，我们不访问最后一个元素，这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，
//否则就无法跳到最后一个位置了。如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次「不必要的跳跃次数」，因此我们不必访问最后一个元素。

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/jump-game-ii/solution/tiao-yue-you-xi-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func jump(nums []int) int {
	step := 0
	maxReach := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		maxReach = max(maxReach, i+nums[i])
		if i == end {
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
