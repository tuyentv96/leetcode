class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        freq = [0] * 26
        
        for c in magazine:
            freq[ord(c) - ord('a')]+=1
            
        for c in ransomNote:
            freq[ord(c) - ord('a')]-=1
            if freq[ord(c) - ord('a')]<0:
                return False
        
        return True