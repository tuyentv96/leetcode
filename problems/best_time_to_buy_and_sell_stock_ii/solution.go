func maxProfit(prices []int) int {
	sum := 0
	i := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		low := prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		sum += prices[i] - low
	}

	return sum
}