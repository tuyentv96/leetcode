class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        sub = [nums[0]]

        for num in nums[1:]:
            if num > sub[-1]:
                sub.append(num)
            else:
                # i = self.bs_leftmost(sub, num)
                i = bisect.bisect_left(sub, num)
                sub[i] = num
        return len(sub)

    def bs_leftmost(self, nums: List[int], num: int) -> int:
        lo, hi = 0, len(nums)-1
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if num == mid:
                hi = mid
            elif nums[mid] < num:
                lo = mid +1
            else:
                hi = mid -1

        return lo
