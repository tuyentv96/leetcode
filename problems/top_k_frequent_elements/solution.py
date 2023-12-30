
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for num in nums:
            count[num] = count.get(num, 0) + 1

        res = []
        for num, c in count.items():
            heappush(res, (c, num))
            if len(res) > k:
                heappop(res)

        return [num for (c, num) in res]
        