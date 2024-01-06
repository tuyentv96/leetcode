class Solution:
    def countSubstrings(self, s: str) -> int:
        res = 0
        for i in range(len(s)):
            res += self.find_palindrome(i, i, s)
            res += self.find_palindrome(i, i + 1, s)
        
        return res

    def find_palindrome(self, l, r, s):
        res = 0
        while l >= 0 and r < len(s):
            if s[l] == s[r]:
                res += 1
                l -= 1
                r += 1
            else: break
        return res