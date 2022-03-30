class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        if len(arr) <= k:
            return arr

        lo, hi = 0, len(arr)
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if arr[mid] < x:
                lo = mid + 1
            else:
                hi = mid
        
        # arr[lo-1] < x < arr[lo]
        left = lo - 1
        right = lo
        
        while right - left - 1 < k:
            if left == -1:
                right += 1
                continue
            
            if right == len(arr) or (arr[right] - x >= x - arr[left]):
                left -= 1
            else:
                right += 1
        
        return arr[left+1:right]
        
        