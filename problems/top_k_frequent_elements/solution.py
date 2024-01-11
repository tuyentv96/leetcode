class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for i in nums:
            count[i] = count.get(i, 0) + 1
        
        counts = list(count.items())
        n = len(counts)
        
        pivot = self.quick_select(0, n - 1, n - k, counts)
        return [num for (num, count) in counts[n-k:]]
        
    def quick_select(self, l, r, target_idx, counts):
        print(l, r, target_idx)
        if l > r:
            return
        
        pivot_idx = self.partition(l, r, counts)
        if target_idx == pivot_idx:
            return
        elif pivot_idx < target_idx:
            self.quick_select(pivot_idx + 1, r, target_idx, counts)
        else:
            self.quick_select(l, pivot_idx - 1, target_idx, counts)

        
    def partition(self, l, r, counts):
        pivot = r
        j = l
        for i in range(l, r):
            if counts[i][1] < counts[pivot][1]:
                counts[i], counts[j] = counts[j], counts[i]
                j += 1
        print(counts)
        counts[pivot], counts[j] = counts[j], counts[pivot]
        print(r, j, counts)
        return j