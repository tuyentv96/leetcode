class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        def reverse(left: int, right: int):
            while left < right:
                nums[left], nums[right] = nums[right], nums[left]
                left+=1
                right-=1
        n = len(nums)
        k = k % n
        # reverse first half
        reverse(0, n - k - 1)
        reverse(n - k, n - 1)
        reverse(0, n - 1)
        