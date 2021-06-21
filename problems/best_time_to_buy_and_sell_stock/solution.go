func maxProfit(prices []int) int {
	max := 0
	min := 1 << 32
	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		} else if max < (prices[i] - min) {
			max = prices[i] - min
		}
	}

	return max
}
