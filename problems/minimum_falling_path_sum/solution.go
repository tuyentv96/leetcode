func minFallingPathSum(matrix [][]int) int {
    n:=len(matrix)
    dp:=make([][]int,n)
    for i:=0;i<n;i++{
        dp[i]=make([]int,n)
        dp[0][i]=matrix[0][i]
    }
    
    for row:=1;row<n;row++{
        for col:=0;col<n;col++{
            if col==0{
                dp[row][col]=min(dp[row-1][col],dp[row-1][col+1])+matrix[row][col]
                continue
            }
            
            if col==n-1{
                dp[row][col]=min(dp[row-1][col],dp[row-1][col-1])+matrix[row][col]
                continue
            }
            
            dp[row][col]=min(dp[row-1][col-1],min(dp[row-1][col],dp[row-1][col+1]))+matrix[row][col]
        }
    }
    
    res:=dp[n-1][0]
    for i:=1;i<n;i++{
        res=min(res,dp[n-1][i])
    }
    
    return res
}

func min(a,b int) int{
    if a<b{
        return a
    }
    
    return b
}