class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        freq1 = [0]*26
        freq2 = [0]*26

        left, right = 0, 0
        
        def equal():
            for i in range(26):
                if freq1[i] != freq2[i]:
                    return False
            return True
        
        def enough():
            for i in range(26):
                if freq1[i] > 0 and freq1[i] > freq2[i]:
                        return False
            return True
        
        for c in s1:
            freq1[ord(c) - ord('a')]+=1
        
        while right < len(s2):
            freq2[ord(s2[right]) - ord('a')] += 1
            while enough() and left <= right:
                if equal():
                    return True
                freq2[ord(s2[left]) - ord('a')] -= 1
                left+=1
            
            right+=1
        
        return False
            
                
            
            