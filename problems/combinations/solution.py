class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        def backtrack(i: int, cur: List[int]):
            if len(cur) == k:
                result.append(cur.copy())
                
            for i in range(i, n+1):
                cur.append(i)
                backtrack(i+1, cur)
                cur.pop()
                
        result = []
        backtrack(1, [])
        return result