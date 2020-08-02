func maxArea(height []int) int {
	var left, right = 0, len(height) - 1
	var max = 0
	for left < right {
		var y int
		if height[left] < height[right] {
			y = height[left]
		} else {
			y = height[right]
		}

		if y*(right-left) > max {
			max = y * (right - left)
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}