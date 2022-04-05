class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        m = {}
        sum = 0
        rs = 0
        m[0] = -1
        for i, num in enumerate(nums):
            sum += 1 if num == 1 else -1
            print("sum", sum)
            print("get sum", m.get(sum, None))
            if m.get(sum, None) is not None:
                rs = max(rs, i - m[sum])
                print("msum", m[sum])
            else:
                m[sum] = i
        
        return rs