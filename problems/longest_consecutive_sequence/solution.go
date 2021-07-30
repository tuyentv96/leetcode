func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	var result int
	for _, num := range nums {
		if _, ok := m[num-1]; !ok {
			current := 0
			curNum := num
			for true {
				if v := m[curNum]; !v {
					break
				}

				current++
				curNum++
			}

			result = max(result, current)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
