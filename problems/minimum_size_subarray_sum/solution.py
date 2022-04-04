class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        sum = 0
        i = 0
        left = 0
        result = float('inf')
        
        for i in range(len(nums)):
            sum += nums[i]
            while sum >= target:
                result = min(result, i - left +1)
                sum -= nums[left]
                left += 1
            
        return result if result != float('inf') else 0
                