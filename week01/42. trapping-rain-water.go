package leetcode

// 对于位置left而言，它左边最大值一定是left_max，右边最大值“大于等于”right_max，这时候，
//如果left_max<right_max成立，那么它就知道自己能存多少水了。
//无论右边将来会不会出现更大的right_max，都不影响这个结果。
//所以当left_max<right_max时，我们就希望去处理left下标，反之，我们希望去处理right下标。
func trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left <= right {
		if leftMax < rightMax {
			ans += max(0, leftMax-height[left])
			leftMax = max(leftMax, height[left])
			left++
		} else {
			ans += max(0, rightMax-height[right])
			rightMax = max(rightMax, height[right])
			right--
		}
	}
	return ans
}
