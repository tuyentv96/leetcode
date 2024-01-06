# dp[i] = max(dp[i-1] + nums[i], nums[i])

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        res = float('-inf')
        cur = float('-inf')
        for i in range(len(nums)):
            cur = max(cur + nums[i], nums[i])
            res = max(res, cur)
        
        return res