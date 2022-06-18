class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        left, right = 0, n-1
        rs = [0] * n
        
        for i in range(n-1, -1, -1):
            num = 0
            if abs(nums[left]) > abs(nums[right]):
                num = nums[left]
                left+=1
            else:
                num = nums[right]
                right-=1
            
            rs[i] = num * num
            
        return rs
                