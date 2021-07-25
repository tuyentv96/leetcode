func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1
	for left <= right {
		mid := left + (right-left)/2
		r := mid / cols
		c := mid % cols
		v := matrix[r][c]
		if target == v {
			return true
		}

		if target < v {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
