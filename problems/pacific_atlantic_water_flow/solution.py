class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        ROWS = len(heights)
        COLS = len(heights[0])
        visited_pac = set()
        visited_atl = set()

        def dfs(r, c, visited, prev_height):
            if c < 0 or r < 0 or c >= COLS or r >= ROWS or (r,c) in visited or heights[r][c] < prev_height:
                return

            visited.add((r, c))
            dfs(r+1, c, visited, heights[r][c])
            dfs(r-1, c, visited, heights[r][c])
            dfs(r, c+1, visited, heights[r][c])
            dfs(r, c-1, visited, heights[r][c])

        for c in range(COLS):
            dfs(0, c, visited_pac, heights[0][c])
            dfs(ROWS-1, c, visited_atl, heights[ROWS-1][c])

        for r in range(ROWS):
            dfs(r, 0, visited_pac, heights[r][0])
            dfs(r, COLS-1, visited_atl, heights[r][COLS-1])

        return list(visited_pac.intersection(visited_atl))