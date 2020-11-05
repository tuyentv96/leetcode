func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k, l, r int) int {
	if l == r {
		return nums[l]
	}

	if k > 0 && k <= r-l+1 {
		mid := partition(nums, l, r)
		if mid-l == k-1 {
			return nums[mid]
		}

		if mid-l > k-1 {
			return quickSelect(nums, k, l, mid-1)
		} else {
			return quickSelect(nums, k-mid+l-1, mid+1, r)
		}
	}

	return 0
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if nums[j] > pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[high], nums[i+1] = nums[i+1], nums[high]

	return i + 1
}
