func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp_i2 := nums[0]
	dp_i1 := Max(nums[0], nums[1])
	if len(nums) == 2 {
		return dp_i1
	}

	dp_i := Max(dp_i1, dp_i2+nums[2])
	for i := 2; i < len(nums); i++ {
		dp_i = Max(dp_i1, dp_i2+nums[i])
		dp_i2 = dp_i1
		dp_i1 = dp_i
	}

	return dp_i
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
