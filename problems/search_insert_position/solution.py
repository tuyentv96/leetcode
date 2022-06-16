class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        if len(nums) == 0:
            return 0

        lo, hi = 0, len(nums)-1
        
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            print(lo, mid, hi)

            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                hi = mid -1
            else:
                lo = mid + 1
        
        return lo