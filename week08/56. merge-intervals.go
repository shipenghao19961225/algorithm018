package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
			continue
		}
		res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
	}
	return res

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
