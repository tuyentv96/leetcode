func removeDuplicates(nums []int) int {
	var begin int
	for i := 1; i < len(nums); i++ {
		if begin == 0 || nums[i] != nums[begin] || nums[i] != nums[begin-1] {
			begin++
			nums[begin], nums[i] = nums[i], nums[begin]
		}
	}

	return begin + 1
}