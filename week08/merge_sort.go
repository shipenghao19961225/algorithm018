package leetcode

import "fmt"

func main() {
	array := []int{2, 3, 1, 4}
	mergeSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func mergeSort(array []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	mergeSort(array, left, mid)
	mergeSort(array, mid+1, right)
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
