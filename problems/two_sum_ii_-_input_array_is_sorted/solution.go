func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		j, ok := m[c]
		if ok && i != j {
			return []int{i+1, j+1}
		}
	}

	return []int{0, 0}
}