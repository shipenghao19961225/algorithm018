package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	low, high := 0, row*col-1
	for low <= high {
		mid := low + (high-low)/2
		value := matrix[mid/col][mid%col]
		if value == target {
			return true
		} else if value > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
