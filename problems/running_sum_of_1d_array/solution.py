class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        result = [0] * len(nums)
        
        prev = 0
        for i in range(len(nums)):
            result[i] = prev + nums[i]
            prev = result[i]
            
        return result