class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        left, right = self.bisect_left(nums, target), self.bisect_right(nums, target) - 1
        return [
            left if left < len(nums) and nums[left] == target else -1,
            right if 0 <= right < len(nums) and nums[right] == target else -1
        ]
      
    def bisect_left(self, a, x):
        lo, hi = 0, len(a)
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if a[mid] < x:
                lo = mid + 1
            else: hi = mid
        return lo
            
    def bisect_right(self, a, x):
        lo, hi = 0, len(a)
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if a[mid] <= x:
                lo = mid + 1
            else: hi = mid
            
        return lo
            