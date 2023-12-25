class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0

        cur = nums[0]
        result = cur
        
        for i in range(1, len(nums)):
            num = nums[i]
            cur = max(cur + num, num)
            result = max(result, cur)

        return result