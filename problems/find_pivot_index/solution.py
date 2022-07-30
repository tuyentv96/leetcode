class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        s = sum(nums)
        prefix_sum = 0
        for i, num in enumerate(nums):
            if prefix_sum == (s - prefix_sum - num):
                return i
            
            prefix_sum += num
            
        return -1