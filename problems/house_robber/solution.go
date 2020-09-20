func rob(nums []int) int {
	var prev int
	var cur int
	var tmp int

	for i := 0; i < len(nums); i++ {
		tmp = cur
		cur = max(cur, prev+nums[i])
		prev = tmp
	}

	return cur
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}