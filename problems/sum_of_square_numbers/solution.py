class Solution:
    def judgeSquareSum(self, c: int) -> bool:
        def search(s: int) -> bool:
            left = 0
            right = s
            while left <= right:
                mid = left + (right - left) // 2
                if mid*mid == s:
                    return True
                elif mid*mid < s:
                    left = mid + 1
                else: right = mid - 1
            return False
            
        for num in range(0, (c+1)):
            if num * num > c:
                return False

            if search(c - num*num):
                return True
        
        return False