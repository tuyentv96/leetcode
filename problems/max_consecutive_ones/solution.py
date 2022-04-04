class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        i=0
        n=len(nums)
        result = 0
            
        while i < n:
            while i < n and nums[i] == 0:
                i+=1
            
            j = 0
            while i < n and nums[i] == 1:
                j+=1
                i+=1
            
            result = max(result, j)
        
        return result