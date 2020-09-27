func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func canJump(nums []int) bool {
	var maxReach int
	for i := 0; i < len(nums); i++ {
		if maxReach < i {
			return false
		}

		maxReach = max(maxReach, i+nums[i])
	}

	return true
}