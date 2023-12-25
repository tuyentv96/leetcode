class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        freq = [0 for i in range(26)]

        for c in s:
            freq[ord(c) - 97] += 1

        for c in t:
            freq[ord(c) - 97] -= 1

        for count in freq:
            if count != 0:
                return False

        return True