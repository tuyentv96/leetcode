class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        def search(m: int) -> int:
            left, right = 0, len(arr) - 1
            
            while left <= right:
                mid = left + (right - left) // 2
                if arr[mid]*2 == m:
                    return mid
                elif (arr[mid])*2 < m:
                    left = mid + 1
                else: right = mid -1
            return -1

        arr.sort()
        
        zero_count = 0
        for i in range(len(arr)):
            if arr[i] != 0:
                j = search(arr[i])
                if j != -1:
                    return True
            else:
                zero_count += 1
                if zero_count == 2:
                    return True
        
        return False
        
        
        