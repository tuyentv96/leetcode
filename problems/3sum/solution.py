# -1,0,1,2,-1,-4
# -1,-1,0,1,2,4
# i = -1 => l=-1,r=0; 

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        n = len(nums)
        for i in range(n):
            l, r = i + 1, n - 1
            target = 0 - nums[i]
            if i > 0 and nums[i] == nums[i-1]:
                continue
                
            while l < r:
                if nums[l] + nums[r] == target:
                    res.append([nums[i], nums[l], nums[r]])
                    l += 1                  
                    while l < r and nums[l-1] == nums[l]:
                        l += 1  
                elif nums[l] + nums[r] < target:
                    l += 1
                else:
                    r -= 1

        return res
