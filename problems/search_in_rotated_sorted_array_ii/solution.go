func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		}

		if nums[m] > nums[l] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[l] {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			l++
		}
	}

	if nums[l] == target {
		return true
	}

	return false
}
