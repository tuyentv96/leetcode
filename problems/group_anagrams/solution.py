class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        def calculate_freq(s: str):
            arr = [0] * 26
            for c in s:
                arr[ord(c) - ord('a')] += 1
            
            return tuple(arr)
        
        result = {}
        for s in strs:
            freq = calculate_freq(s)
            if freq not in result:
                result[freq] = []
            result[freq].append(s)
        
        return result.values()