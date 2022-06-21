class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        if len(nums1) > len(nums2):
            return self.intersect(nums2, nums1)

        nums1.sort()
        nums2.sort()
        
        i1, i2 = 0, 0
        k = 0
        
        while i1 < len(nums1) and i2 < len(nums2):
            if nums1[i1] == nums2[i2]:
                nums1[k] = nums1[i1]
                k+=1
                i1+=1
                i2+=1
            elif nums1[i1] < nums2[i2]:
                i1+=1
            else: i2+=1
        
        return nums1[:k]
        