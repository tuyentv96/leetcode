func isMonotonic(nums []int) bool {
	var inc, desc bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			inc = true
		}

		if nums[i] > nums[i+1] {
			desc = true
		}
	}

	if inc && desc {
		return false
	}

	return true
}