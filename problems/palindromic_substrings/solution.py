class Solution:
    def countSubstrings(self, s: str) -> int:
        result = 0

        for i in range(len(s)):
            result += self.isPalindrome(s,i,i)

            result += self.isPalindrome(s,i,i+1)

        return result

    def isPalindrome(self,s: str, lo: int, hi: int) -> int:
        result = 0
        while lo>=0 and hi<len(s):
            if s[lo]!=s[hi]:
                return result
            
            lo-=1
            hi+=1
            result += 1

        return result
