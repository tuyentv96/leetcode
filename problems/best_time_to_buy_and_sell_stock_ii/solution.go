func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var max int

	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			max += prices[i+1] - prices[i]
		}
	}

	return max
}