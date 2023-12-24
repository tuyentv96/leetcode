class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0

        min_price = prices[0]
        max_profit = 0

        for i in range(1, len(prices)):
            price = prices[i]
            min_price = min(min_price, price)

            max_profit = max(max_profit, price - min_price)
        
        return max_profit

