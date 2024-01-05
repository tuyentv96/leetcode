from collections import defaultdict
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list)

        for str in strs:
            res[self.build_hash(str)].append(str)
        
        return res.values()

    def build_hash(self, s):
        arr = [0] * 26
        for c in s:
            arr[ord(c) - ord('a')] += 1
            
        return tuple(arr)