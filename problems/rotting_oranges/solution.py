class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        if len(grid)==0:
            return 0

        m = len(grid)
        n = len(grid[0])
        queue = collections.deque()
        
        for i in range(m):
            for j in range(n):
                if grid[i][j]==2:
                    queue.append((i, j))
        
        result = -1
        dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        while queue:
            total = len(queue)
            result += 1
            for _ in range(total):
                i, j = queue.popleft()
                
                for r, c in dirs:
                    r, c = r+i, c+j
                    if r < 0 or c < 0 or r >=m or c >=n or grid[r][c]!=1:
                        continue
                    
                    grid[r][c]=2
                    queue.append((r, c))
                    
        for i in range(m):
            for j in range(n):
                if grid[i][j]==1:
                    return -1
        
        return result if result!=-1 else 0