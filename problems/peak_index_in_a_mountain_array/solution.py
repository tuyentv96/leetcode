class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        if len(arr) < 3:
            return -1

        lo = 0
        hi = len(arr) - 1
        
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if arr[mid] > arr[mid+1] and arr[mid-1] < arr[mid]: 
                return mid
            elif arr[mid] < arr[mid+1]:
                lo = mid + 1
            else:
                hi = mid - 1
        
        return -1