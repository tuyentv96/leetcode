func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] += 1
	}

	for i := range nums {
		if m[nums[i]] == 1 {
			return nums[i]
		}
	}

	return 0
}
