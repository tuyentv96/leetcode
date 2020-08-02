func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}

		if target < nums[i] {
			return i
		}
		
		if i == len(nums)-1 && target > nums[i] {
			return i + 1
		}

		if target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}

	return -1
}