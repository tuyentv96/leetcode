class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        if len(mat) == 0:
            return []
        m = len(mat)
        n = len(mat[0])
        
        result = [[-1]*n for _ in range(m)]
        
        queue = collections.deque()
        for i in range(m):
            for j in range(n):
                if mat[i][j]==0:
                    result[i][j]=0
                    queue.append((i, j))
                    
        dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        while queue:
            i, j = queue.popleft()
            
            for r, c in dirs:
                row, col = i+r, j+c
                if row < 0 or col < 0 or row >= m or col >=n or result[row][col]!=-1:
                    continue
                
                result[row][col] = result[i][j]+1
                queue.append((row, col))
            
        return result
        
            
