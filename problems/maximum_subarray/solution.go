func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var maxSoFar, max = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+max > nums[i] {
			max = max + nums[i]
		} else {
			max = nums[i]
		}

		if max > maxSoFar {
			maxSoFar = max
		}
	}

	return maxSoFar
}
