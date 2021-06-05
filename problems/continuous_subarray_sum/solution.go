func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	r := 0
	for i := 0; i < len(nums); i++ {
		r += nums[i]
		if k != 0 {
			r = r % k
		}

		if _, ok := m[r]; ok {
			if i-m[r] >= 2 {
				return true
			}
		} else {
			m[r] = i
		}
	}

	return false
}
