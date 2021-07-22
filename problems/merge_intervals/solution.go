func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	current := intervals[0]
	result = append(result, current)

	for i := range intervals {
		currentEnd := current[1]
		nextBegin := intervals[i][0]
		nextEnd := intervals[i][1]

		if currentEnd >= nextBegin {
			current[1] = max(currentEnd, nextEnd)
		} else {
			current = intervals[i]
			result = append(result, current)
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
