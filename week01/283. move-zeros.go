package leetcode

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 这里的交换操作比较精髓
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
