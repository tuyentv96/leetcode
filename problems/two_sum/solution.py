class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = dict()
        result = []
        for i, num in enumerate(nums):
            if (target - num) in seen:
                result.extend((i, seen[target - num]))

            if num not in seen:
                seen[num]=i
        
        return result