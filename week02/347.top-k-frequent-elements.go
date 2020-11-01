package leetcode

import "sort"

// 利用库的排序
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	s := make([]int, 0, len(nums))
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			m[num]++
		} else {
			m[num] = 1
			s = append(s, num)
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return m[s[i]] > m[s[j]]
	})
	return s[:k]
}
