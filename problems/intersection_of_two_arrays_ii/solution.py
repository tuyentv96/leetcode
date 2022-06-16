class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        if len(nums1) > len(nums2):
            return self.intersect(nums2, nums1)

        m = dict()
        for num in nums1:
            m[num] = m.get(num, 0) + 1
        
        k = 0
        for num in nums2:
            cnt = m.get(num, 0)
            if cnt > 0:
                nums1[k] = num
                k += 1
                m[num] = cnt - 1
        
        return nums1[:k]
        