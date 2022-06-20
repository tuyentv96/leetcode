class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        left, right = 0, len(numbers) - 1
        
        while left <= right:
            su = numbers[left] + numbers[right]
            if su == target:
                return [left+1, right+1]
            elif su < target:
                left += 1
            else: right -= 1
        
        return []