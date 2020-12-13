package leetcode

var count int

func reversePairs(nums []int) int {
	count = 0
	mergeSort(nums, 0, len(nums)-1)
	return count
}
func mergeSort(array []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	mergeSort(array, left, mid)
	mergeSort(array, mid+1, right)
	i, j := left, mid+1
	for i <= mid && j <= right {
		if array[i] > array[j]*2 {
			count += mid - i + 1
			j++
		} else {
			i++
		}
	}
	merge(array, left, mid, right)
}

func merge(array []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j := left, mid+1
	k := 0
	for i <= mid && j <= right {
		if array[i] <= array[j] {
			tmp[k] = array[i]
			i++
			k++
		} else {
			tmp[k] = array[j]
			j++
			k++
		}
	}
	for i <= mid {
		tmp[k] = array[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = array[j]
		j++
		k++
	}
	copy(array[left:right+1], tmp)
}
