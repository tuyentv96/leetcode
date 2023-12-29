class Solution:
    def count_palindrome(self, s, l, r) -> int:
        res = 0
        while l >= 0 and r < len(s) and s[l] == s[r]:
            res += 1
            r += 1
            l -= 1
        
        return res

    def countSubstrings(self, s: str) -> int:
        res = 0
        for i in range(len(s)):
            res += self.count_palindrome(s, i, i)
            res += self.count_palindrome(s, i, i+1)
        
        return res
