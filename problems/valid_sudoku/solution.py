class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = 9
        cols = 9

        rows_seen = [set() for _ in range(rows)]
        cols_seen = [set() for _ in range(cols)]
        grid_seen = [set() for _ in range(cols)]

        for i in range(rows):
            for j in range(cols):
                val = board[i][j]
                if val != ".":
                    if val in rows_seen[i]:
                        return False
                    rows_seen[i].add(val)

                    if val in cols_seen[j]:
                        return False                
                    cols_seen[j].add(val)
                    
                    grid_idx = (i // 3) * 3 + j // 3
                    if val in grid_seen[grid_idx]:
                        return False
                    grid_seen[grid_idx].add(val)
                
        return True