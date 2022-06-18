class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        last_non_zero_index = 0
        
        for i in range(n):
            if nums[i] != 0:
                nums[last_non_zero_index] = nums[i]
                last_non_zero_index+=1
        
        print(last_non_zero_index)
        for i in range(last_non_zero_index, n):
            nums[i] = 0
        
        return
            