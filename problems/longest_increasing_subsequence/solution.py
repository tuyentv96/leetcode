class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        sub = [nums[0]]

        for num in nums[1:]:
            i = self.bs_leftmost(sub, num)
            if num > sub[-1]:
                sub.append(num)
            else:
                sub[i] = num
        return len(sub)

    def bs_leftmost(self, nums: List[int], num: int) -> int:
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if nums[mid] < num:
                lo = mid +1
            else:
                hi = mid

        return lo
