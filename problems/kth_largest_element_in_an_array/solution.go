func findKthLargest(nums []int, k int) int {
	n := len(nums)
	i := n - k
	l, r := 0, len(nums)-1
	for l < r {
		pivot := partition(nums, l, r)
		if pivot == i {
			return nums[pivot]
		}

		if pivot > i {
			r = pivot - 1
		} else {
			l = pivot + 1
		}
	}

	return nums[l]
}

func partition(nums []int, left, right int) int {
	pivot := right
	j := left
	for i := left; i < right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	nums[pivot], nums[j] = nums[j], nums[pivot]
	return j
}
