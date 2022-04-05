class Solution:
    def maxSubArrayLen(self, nums: List[int], k: int) -> int:
        m = {}
        rs = 0
        sum = 0
        
        # m[0] = -1
        
        for i, num in enumerate(nums):
            sum += num
            
            if sum == k:
                rs = i + 1
            
            if sum - k in m:
                rs = max(rs, i - m[sum - k])
            
            if sum not in m:
                m[sum] = i
        
        return rs