class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if len(matrix) < 1:
            return False
        
        total_row = len(matrix)
        total_col = len(matrix[0])

        left = 0
        right = total_row * total_col - 1
        
        while left <= right:
            mid = left + (right - left) // 2
            row = mid // total_col
            col = mid % total_col
                        
            if matrix[row][col] == target:
                return True
            elif matrix[row][col] < target:
                left = mid + 1
            else: right = mid - 1
        
        return False