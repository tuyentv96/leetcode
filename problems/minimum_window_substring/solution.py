class Solution:
    def minWindow(self, s: str, t: str) -> str:
        have = 0
        need = 0
        res = None
        count_s = {}
        count_t = {}
        min_size = float('Infinity')

        for c in t:
            count_t[c] = count_t.get(c, 0) + 1
        
        need = len(count_t)
        l = 0
        for r in range(len(s)):
            c = s[r]
            count_s[c] = count_s.get(c, 0) + 1
            if count_s[c] == count_t.get(c, 0):
                have += 1
    
            while have == need:
                if r - l < min_size:
                    min_size = r - l
                    res = (l, r)

                count_s[s[l]] -= 1
                if count_s[s[l]] < count_t.get(s[l], 0):
                    have -= 1

                l += 1
            

        return s[res[0]:res[1]+1] if res else ""
                