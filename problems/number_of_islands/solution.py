class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        res = 0

        def dfs(grid: List[List[str]], i, j):
            if i < 0 or i >= len(grid) or j < 0 or j >= len(grid[0]):
                return

            if grid[i][j] == '1':
                grid[i][j] = 'x'
                dfs(grid, i + 1, j)
                dfs(grid, i - 1, j)
                dfs(grid, i, j + 1)
                dfs(grid, i, j - 1)



        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == '1':
                    res += 1
                    dfs(grid, i, j)
        
        return res
    