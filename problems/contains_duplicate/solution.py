class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        arr = set()

        for num in nums:
            if num in arr:
                return True

            arr.add(num)

        return False