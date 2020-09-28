func trap(heights []int) int {
	if len(heights) <= 2 {
		return 0
	}

	var maxLeft, maxRight = heights[0], heights[len(heights)-1]
	var left, right = 0, len(heights) - 1
	var water int
	for left < right {
		if heights[left] < heights[right] {
			left++
			maxLeft = max(maxLeft, heights[left])
			water += maxLeft - heights[left]
		} else {
			right--
			maxRight = max(maxRight, heights[right])
			water += maxRight - heights[right]
		}
	}

	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}