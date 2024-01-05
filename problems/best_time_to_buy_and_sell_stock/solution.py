class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        min_p = prices[0]
        res = 0
        for i in range(1, len(prices)):
            res = max(res, prices[i] - min_p)
            min_p = min(min_p, prices[i])
        return res