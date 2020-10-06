func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}

func rob(nums []int) int {
    n:=len(nums)
    dp:=make([]int,n)
    
    if n==0{
        return 0
    }
    
    if n>0{
        dp[0]=nums[0]
    }
    
    if n>1{
        dp[1]=max(dp[0],nums[1])
    }
    
    for i:=2;i<n;i++{
        dp[i]=max(dp[i-2]+nums[i],dp[i-1])
    }
    
    return dp[n-1]
}