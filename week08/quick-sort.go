package leetcode

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(array, left, right)
	quickSort(array, left, pivot-1)
	quickSort(array, pivot+1, right)
}
func partition(array []int, left, right int) int {
	val := array[right]
	counter := left
	for i := left; i < right; i++ {
		if array[i] < val {
			array[i], array[counter] = array[counter], array[i]
			counter++
		}
	}
	array[right], array[counter] = array[counter], array[right]
	return counter
}
