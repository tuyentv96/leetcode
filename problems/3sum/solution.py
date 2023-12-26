class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        sorted_nums = sorted(nums)
        length = len(nums)
        print(sorted_nums)

        result = []
        for i in range(length):
            left = i + 1
            right = length - 1

            if i > 0 and sorted_nums[i] == sorted_nums[i - 1]:
                continue

            while left < right:
                sum = sorted_nums[i] + sorted_nums[left] + sorted_nums[right]
                if sum == 0:
                    result.append([sorted_nums[i], sorted_nums[left], sorted_nums[right]])

                    while left < right and sorted_nums[left] == sorted_nums[left + 1]:
                        left += 1
                    left += 1
                elif sum < 0:
                    left += 1
                else:
                    right -= 1

        return result