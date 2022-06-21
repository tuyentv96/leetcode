class Solution:
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        result = [[0]*c for _ in range(r)]
        mat_row = 0
        mat_col = 0
        m = len(mat)
        n = len(mat[0])
        
        if r * c != m * n:
            return mat
        
        for row in range(r):
            for col in range(c):
                if mat_col == n:
                    mat_row += 1
                    mat_col = 0
                
                result[row][col] = mat[mat_row][mat_col]
                mat_col += 1

        return result