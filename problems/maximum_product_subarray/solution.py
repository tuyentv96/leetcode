class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        max_p = 1
        min_p = 1
        res = float('-inf')

        for num in nums:
            if num < 0:
                min_p, max_p = max_p, min_p
            min_p = min(min_p * num, num)
            max_p = max(max_p * num, num)
            res = max(res, max_p)
        return res
