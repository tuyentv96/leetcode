func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var min = prices[0]
	var max int

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - min
		if diff > max {
			max = diff
		}

		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}