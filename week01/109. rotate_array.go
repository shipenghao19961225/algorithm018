package leetcode

// 可以类比同学换座位
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	// 如果正好可以整除，就直接返回，否则会导致数组越界
	if k == 0 || n <= 1 {
		return
	}
	count := 0
	start := 0
	tmp := nums[0]
	for i := 0; count < n; {
		i = (i + k) % n
		tmp, nums[i] = nums[i], tmp
		count++
		// 如果回到了起点,需要从下一个位置，继续交换，直到count = k
		if i == start {
			i++
			start++
			tmp = nums[i]
		}
	}
}
