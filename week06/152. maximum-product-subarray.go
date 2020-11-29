package leetcode

func maxProduct(nums []int) int {
	minVal, maxVal := nums[0], nums[0]
	res := maxVal
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal

		}
		minVal = min(minVal*nums[i], nums[i])
		maxVal = max(maxVal*nums[i], nums[i])
		res = max(res, maxVal)
	}

	return res

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
