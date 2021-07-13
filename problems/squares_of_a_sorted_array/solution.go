func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	index := len(nums) - 1
	result := make([]int, len(nums))
	for l <= r {
		if abs(nums[l]) > abs(nums[r]) {
			result[index] = nums[l] * nums[l]
			l++
		} else {
			result[index] = nums[r] * nums[r]
			r--
		}
		index--
	}

	return result
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
