# ADOBECODEBANC

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        have = 0
        s_count = {}
        t_count = {}

        for c in t:
            t_count[c] = t_count.get(c, 0) + 1

        l = 0
        res = None
        min_size = float('inf')
        need = len(t_count)
        for r in range(len(s)):
            c = s[r]
            s_count[c] = s_count.get(c, 0) + 1
            if s_count[c] == t_count.get(c, 0):
                have += 1
            
            while have == need:
                if r - l < min_size:
                    res = (l, r)
                    min_size = r - l

                s_count[s[l]] -= 1
                if s_count[s[l]] < t_count.get(s[l], 0):
                    have -= 1
                l += 1

        return s[res[0]:res[1] + 1] if res else ""