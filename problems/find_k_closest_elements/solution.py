class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        if len(arr) <= k:
            return arr

        lo, hi = 0, len(arr)
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if arr[mid] < x:
                lo = mid + 1
            else: hi = mid
        
        left, right = lo-1, lo
        print(left, right)
        
        while right - left - 1 < k:
            print("db", left, right)
            if left == -1:
                right+=1
                continue
            if right == len(arr):
                left-=1
                continue
                
            if abs(arr[left] - x) <= abs(arr[right] - x):
                left-=1
            else: right+=1
        
        print("ans", left+1, right)

        return arr[left+1:right]
        
        