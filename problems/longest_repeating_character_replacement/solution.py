class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        char_count = {}
        l = 0
        r = 0
        res = 0

        while r < len(s):
            char_count[s[r]] = char_count.get(s[r], 0) + 1

            while (r - l + 1) - max(char_count.values()) > k:
                char_count[s[l]] -= 1
                l += 1

            res = max(res, r - l + 1)
            r += 1
        
        return res