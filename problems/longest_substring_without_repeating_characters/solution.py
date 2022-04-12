class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        m = [0 for i in range(256)]
        left = 0
        i = 0
        rs = 0
        
        while i < len(s):
            j = ord(s[i])
            m[j] += 1
            
            while m[j] > 1:
                m[ord(s[left])] -= 1
                left += 1

            rs = max(rs, i - left + 1)
            i += 1
        
        return rs