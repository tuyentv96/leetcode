func maxSubArray(nums []int) int {
	res := -(1 << 31)
	sum := 0
	for i := range nums {
		sum = max(nums[i], sum+nums[i])
		res = max(res, sum)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
