class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for num in nums:
            count[num] = count.get(num, 0) + 1

        counts = list(count.items())
        n = len(counts)
        self.quick_select(counts, n-k, 0, n-1)
        return [num for (num, c) in counts[n-k:]]

    def partition(self, nums, left, right):
        pivot = right
        j = left
        for i in range(left, right):
            if nums[i][1] < nums[pivot][1]:
                nums[i], nums[j] = nums[j], nums[i]
                j += 1
        
        nums[pivot], nums[j] = nums[j], nums[pivot]
        return j

    def quick_select(self, nums, index, left, right):
        if left > right:
            return
        
        pivot = self.partition(nums, left, right)
        if pivot == index:
            return
        elif pivot < index:
            self.quick_select(nums, index, pivot + 1, right)
        else:
            self.quick_select(nums, index, left, pivot - 1)
  