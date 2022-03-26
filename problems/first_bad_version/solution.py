# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        lo, hi = 1, n
        
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if isBadVersion(mid) != True:
                lo = mid + 1
            else:
                hi = mid
        
        return lo
        