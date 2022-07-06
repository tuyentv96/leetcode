class Solution:
    def chalkReplacer(self, chalk: List[int], k: int) -> int:
        total = k % sum(chalk)
        for i, v in enumerate(chalk):
            if total < v:
                return i
            
            total -= v
        
        return 0