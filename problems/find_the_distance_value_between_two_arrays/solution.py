class Solution:
    def findTheDistanceValue(self, arr1: List[int], arr2: List[int], d: int) -> int:
        rs = 0
        arr2=sorted(arr2)
        print(arr1)
        print(arr2)
        for v1 in arr1:
            keep = True
            begin = v1 - d
            end = v1 + d
            lo, hi = 0, len(arr2)-1
            while lo <= hi:
                mid = lo + (hi - lo) // 2
                if arr2[mid] >= begin and arr2[mid]<=end:
                    keep = False
                    break
                elif arr2[mid] < begin:
                    lo = mid + 1
                else: hi = mid - 1
            if keep:
                rs += 1
        return rs