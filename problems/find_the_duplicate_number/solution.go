func findDuplicate(nums []int) int {
	var i, j = 0, 0
	for true {
		i = nums[i]
		j = nums[nums[j]]
		if i == j {
			break
		}
	}

	i = 0
	for i != j {
		i = nums[i]
		j = nums[j]
	}

	return j
}