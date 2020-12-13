package leetcode

import "sort"

// 计数排序
func relativeSortArray(arr1 []int, arr2 []int) []int {
	res := make([]int, 0, len(arr1))
	bucket := make([]int, 1001)
	for _, num1 := range arr1 {
		bucket[num1]++
	}
	for _, num2 := range arr2 {
		for bucket[num2] > 0 {
			res = append(res, num2)
			bucket[num2]--
		}
	}
	for i, _ := range bucket {
		for bucket[i] > 0 {
			res = append(res, i)
			bucket[i]--
		}
	}
	return res
}

// 库函数排序

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for i, num2 := range arr2 {
		m[num2] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		a, b := arr1[i], arr1[j]
		indexA, hasA := m[a]
		indexB, hasB := m[b]
		if hasA && hasB {
			return indexA < indexB
		}
		if hasA || hasB {
			return hasA
		}
		return a < b
	})
	return arr1
}
