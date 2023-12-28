class Solution:
    def maxArea(self, height: List[int]) -> int:
        if len(height) < 2:
            return 0

        length = len(height)
        left = 0
        right = length - 1
        result = 0

        while left <= right:
            print(left, right)
            result = max(result, min(height[right], height[left]) * (right - left))
            if height[left] <= height[right]:
                left += 1
            else:
                right -= 1
        
        return result