func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var prev = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if prev == nums[i] {
			prev = nums[i]
			nums = append(nums[0:i], nums[i+1:]...)
			continue
		}

		prev = nums[i]
	}
	return len(nums)
}