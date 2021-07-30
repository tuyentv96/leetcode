func eraseOverlapIntervals(intervals [][]int) int {
	cnt := 1
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			cnt++
			end = intervals[i][1]
		}
	}

	return len(intervals) - cnt
}
