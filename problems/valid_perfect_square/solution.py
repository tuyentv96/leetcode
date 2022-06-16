class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        if num == 1:
            return True
        lo = 1
        hi = num // 2
        
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            sq = mid*mid
            print(sq, mid)
            if sq == num:
                return True
            elif sq < num:
                lo = mid + 1
            else: hi = mid - 1
        
        return False