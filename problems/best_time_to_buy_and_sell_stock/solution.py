class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0

        rs = 0
        min_buy = prices[0]
        cur_max = 0
        
        for p in prices[1:]:
            if p - min_buy > 0:
                cur_max = max(cur_max, p - min_buy)
                rs = max(rs, cur_max)
            
            if p < min_buy:
                min_buy = p
        
        return rs
        
        