class Solution:
    def lengthOfLongestSubstringTwoDistinct(self, s: str) -> int:
        if len(s) < 3:
            return len(s)
        freq = collections.defaultdict()
        
        hi = 0
        lo = 0
        result = 2

        while hi < len(s):
            freq[s[hi]] = hi
            hi += 1

            if len(freq) == 3:
                del_idx = min(freq.values())
                del freq[s[del_idx]]
                lo = del_idx + 1
            
            result = max(result, hi - lo)
        
        return result
            