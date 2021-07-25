func subarraySum(nums []int, k int) int {
	result := 0
	sum := 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result += m[sum-k]
		m[sum]++
	}

	return result
}
