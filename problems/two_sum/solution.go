func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if complementIndex, ok := m[complement]; ok {
			return []int{complementIndex, i}
		}

		m[v] = i
	}

	return nil
}