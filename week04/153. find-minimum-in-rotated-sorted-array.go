package leetcode

// 题解解释的很清楚
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/solution/yi-wen-jie-jue-4-dao-sou-suo-xuan-zhuan-pai-xu-s-3/
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		if nums[low] <= nums[high] {
			return nums[low]
		}
		mid := low + (high-low)/2
		if nums[mid] >= nums[low] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return -1
}
