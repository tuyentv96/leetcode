func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	j := 1
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] != nums[i-1] {
				nums[j] = nums[i]
				j++
			}
		}
	}

	return j
}
