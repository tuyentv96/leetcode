func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func jump(nums []int) int {
	var maxReach int
	var lastReach int
	var step int
	for i := 0; i < len(nums); i++ {
		if lastReach < i {
			step++
			lastReach = maxReach
		}

		maxReach = max(maxReach, i+nums[i])
	}

	return step
}