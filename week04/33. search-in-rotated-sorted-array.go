package leetcode

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		val := nums[mid]
		if val == target {
			return mid
		} else if val >= nums[low] {
			if target >= nums[low] && target <= val {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target >= val && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
