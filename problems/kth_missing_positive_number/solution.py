class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        lo, hi = 0, len(arr) - 1
        
        
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if arr[mid] - k < mid + 1:
                lo = mid + 1
            else:
                hi = mid - 1
        
        print(lo, hi)
        return hi + 1 + k