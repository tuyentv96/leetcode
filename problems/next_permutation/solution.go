func nextPermutation(nums []int) {
	var j, m int
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			j = i - 1
			m = i
			break
		}
	}

	if m == 0 && j == 0 {
		reverse(nums)
		return
	}

	for i := n - 1; i > j; i-- {
		if i == n-1 && j == n-2 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			break
		}

		if i == n-1 && nums[i] > nums[j] {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			break
		}

		if nums[j] >= nums[i] && nums[j] < nums[i-1] {
			temp := nums[i-1]
			nums[i-1] = nums[j]
			nums[j] = temp
			break
		}
	}

	reverse(nums[j+1:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
}