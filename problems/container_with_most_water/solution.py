class Solution:
    def maxArea(self, height: List[int]) -> int:
        left, right = 0, len(height)-1
        rs = 0
        
        while left < right:
            hl, hr = height[left], height[right]
            rs = max(rs, (right - left) * min(hl, hr))
            if hl < hr:
                left += 1
            else:
                right -= 1
        
        return rs