func deleteAndEarn(nums []int) int {
	points := make([]int, 10001)

	for i := range nums {
		points[nums[i]] += nums[i]
	}

	cur := 0
	prev := 0
	for i := 0; i <= 10000; i++ {
		temp := cur
		cur = max(cur, prev+points[i])
		prev = temp
	}

	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
