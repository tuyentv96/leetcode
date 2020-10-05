func minSubArrayLen(s int, nums []int) int {
	var sum int
	var min, l, r int
	for r = 0; r < len(nums); r++ {
		sum += nums[r]
		for l <= r && sum >= s {
			tmp := r - l + 1
			if min == 0 {
				min = tmp
			}

			if min > tmp {
				min = tmp
			}

			sum -= nums[l]
			l++
		}
	}

	return min
}