from collections import defaultdict
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        count = defaultdict(lambda: 0)
        l = 0
        res = 0
        for r in range(len(s)):
            count[s[r]] += 1
            while count[s[r]] > 1:
                count[s[l]] -= 1
                l += 1
            res = max(res, r - l + 1)
        
        return res