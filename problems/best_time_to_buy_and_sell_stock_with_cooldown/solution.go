func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var buy, sell int
	buy = -prices[0]
	var prevSell int

	for i := 1; i < len(prices); i++ {
		//buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		buy = max(buy, prevSell-prices[i])

		prevSell = sell

		//sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		sell = max(sell, buy+prices[i])
	}
	//for i := 0; i < len(prices)-1; i++ {
	//	if prev-prices[i] > stock {
	//		stock = stock - prices[i]
	//	}
	//
	//	prev = money
	//
	//	if stock+prices[i] > money {
	//		money = stock + prices[i]
	//	}
	//}

	//return sell[len(sell)-1]
	return sell
}
func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
