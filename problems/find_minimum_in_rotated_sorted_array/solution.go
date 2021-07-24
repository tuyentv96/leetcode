func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[r] > nums[l] {
		return nums[l]
	}

	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}

		if nums[m] < nums[m-1] {
			return nums[m]
		}

		if nums[m] > nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return nums[l]
}
