class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        if len(nums) < 1:
            return 0

        imax = nums[0]
        imin = nums[0]
        result = nums[0]

        for i in range(1, len(nums)):
            if nums[i] < 0:
                imax, imin = imin, imax
            
            imax = max(nums[i], imax * nums[i])
            imin = min(nums[i], imin * nums[i])
            result = max(imax, result)

            
        return result