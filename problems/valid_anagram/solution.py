class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        sf = [0] * 26
        tf = [0] * 26

        for c in s:
            sf[ord(c) - ord('a')] += 1
        for c in t:
            tf[ord(c) - ord('a')] += 1

        return sf == tf