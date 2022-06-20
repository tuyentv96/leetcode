class Solution:
    def arrangeCoins(self, n: int) -> int:
        left, right = 1, n
        
        while left <= right:
            mid = left + (right -left) //2
            cur = mid * (mid + 1) //2
            if cur == n:
                return mid
            elif cur < n:
                left = mid + 1
            else: right = mid - 1
        
        return right