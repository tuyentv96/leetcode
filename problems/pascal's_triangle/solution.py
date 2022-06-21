class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        result = [[1]]
        i = 2

        while i <= numRows:
            cur = [0]*i
            for j in range(i):
                tmp = 0
                if j>0:
                    tmp += result[i-2][j-1]
                if j<i-1:
                    tmp += result[i-2][j]
                    
                cur[j] = tmp
            result.append(cur)
            i+=1
        
        return result