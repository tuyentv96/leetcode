class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        visited_p = set()
        visited_a = set()

        ROWS = len(heights)
        COLS = len(heights[0])

        def dfs(r, c, visited, prev):
            if c < 0 or r < 0 or c >= COLS or r >= ROWS or (r, c) in visited or heights[r][c] < prev:
                return

            visited.add((r, c))
            dfs(r + 1, c, visited, heights[r][c])
            dfs(r - 1, c, visited, heights[r][c])
            dfs(r, c + 1, visited, heights[r][c])
            dfs(r, c - 1, visited, heights[r][c])

        for i in range(ROWS):
            dfs(i, 0, visited_p, heights[i][0])
            dfs(i, COLS - 1, visited_a, heights[i][COLS - 1])

        for i in range(COLS):
            dfs(0, i, visited_p, heights[0][i])
            dfs(ROWS - 1, i, visited_a, heights[ROWS - 1][i]) 
            
        return list(visited_p.intersection(visited_a))
