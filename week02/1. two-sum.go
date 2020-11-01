package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		} else {
			m[num] = i
		}
	}
	return []int{}
}
