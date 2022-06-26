class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        if len(grid) == 0:
            return 0

        m, n = len(grid), len(grid[0])
        self.visited = [[False]*n for _ in range(m)]
        self.count = 0
        result = 0
        def dfs(i: int, j: int):            
            if i < 0 or j < 0 or i >= m or j >= n:
                return
            
            if self.visited[i][j]:
                return
            
            if grid[i][j] == 0:
                return
            
            self.count += 1
            self.visited[i][j]=True
            
            dfs(i + 1, j)
            dfs(i - 1, j)
            dfs(i, j + 1)
            dfs(i, j - 1)
                        
        for i in range(m):
            for j in range(n):
                if not self.visited[i][j]:
                    self.count = 0
                    dfs(i, j)
                    result = max(result, self.count)
        
        return result
    
    