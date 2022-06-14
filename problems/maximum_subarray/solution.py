class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        lo = 1
        max_so_far = nums[0]
        cur = nums[0]
        
        while lo < len(nums):
            cur = max(cur + nums[lo], nums[lo])
            max_so_far = max(cur, max_so_far)
            lo += 1
        
        return max_so_far