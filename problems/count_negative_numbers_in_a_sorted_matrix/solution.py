class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        def binary_search(arr: List[int]) -> int:
            left, right = 0, len(arr)
            while left < right:
                mid = left + (right - left) // 2
                if arr[mid] >= 0:
                    left = mid + 1
                else: right = mid
                    
            return left
        
        result = 0
        for arr in grid:
            i = binary_search(arr)
            result += len(arr) - i
        
        return result