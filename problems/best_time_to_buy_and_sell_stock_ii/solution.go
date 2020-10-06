
func maxProfit(prices []int) int {
	var result int
	var buy, sell int
	for sell < len(prices) {
		for sell < len(prices)-1 && prices[sell] < prices[sell+1] {
			sell++
		}

		for buy < sell && prices[buy] > prices[buy+1] {
			buy++
		}

		result += prices[sell] - prices[buy]
		sell++
		buy = sell
	}

	return result
}