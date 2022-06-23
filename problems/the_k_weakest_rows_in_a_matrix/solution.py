class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        def search(arr: List[int]) -> int:
            left, right = 0, len(arr)
            while left < right:
                mid = left + (right - left) // 2
                if arr[mid] == 1:
                    left = mid + 1
                else: right = mid
            return left

        result = []
        for i, row in enumerate(mat):
            total = search(row)
            result.append((total, i))
                
        result.sort(key = lambda x : x[0])
        
        return [i for (total, i) in result[:k]]