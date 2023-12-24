class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        arr = {}

        for idx, num in enumerate(nums):
            if (target-num) in arr:
                return [arr[target-num], idx] 
            
            arr[num] = idx
        
        return []
        