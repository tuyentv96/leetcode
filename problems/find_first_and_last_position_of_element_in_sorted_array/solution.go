func searchRange(nums []int, target int) []int {
	left, right := findLeftIndex(nums, target), findRightIndex(nums, target)
	return []int{left, right}
}

func findLeftIndex(nums []int, target int) int {
	index := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}

		if nums[mid] == target {
			index = mid
		}
	}

	return index
}

func findRightIndex(nums []int, target int) int {
	index := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}

		if nums[mid] == target {
			index = mid
		}
	}

	return index
}