func topKFrequent(nums []int, k int) []int {
	m := buildFreqMap(nums)
	uniqueArray := buildUniqueArray(m)
	n := len(uniqueArray)
	quickSelect(uniqueArray, m, 0, n-1, n-k)
	return uniqueArray[n-k : n]
}

func quickSelect(nums []int, m map[int]int, left, right, index int) int {
	if left < right {
		pivot := partition(nums, m, left, right)
		if pivot == index {
			return pivot
		}

		if index < pivot {
			return quickSelect(nums, m, left, pivot-1, index)
		} else {
			return quickSelect(nums, m, pivot+1, right, index)
		}
	}

	return 0
}

func buildUniqueArray(m map[int]int) []int {
	var result []int
	for k, _ := range m {
		result = append(result, k)
	}

	return result
}

func partition(nums []int, m map[int]int, left, right int) int {
	j := left
	pivot := right
	for i := left; i < right; i++ {
		if m[nums[i]] < m[nums[pivot]] {
			swap(nums, i, j)
			j++
		}
	}

	swap(nums, j, pivot)
	return j
}

func buildFreqMap(nums []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	return m
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
