func rob(nums []int) int {
    n:=len(nums)
    
    if n<2{
        return nums[0]
    }
    
    return max(helper(nums,0,n-2),helper(nums,1,n-1))
}

func helper(nums []int,start, end int) int{
    dp2,dp1,dp:=0,0,0
    
    for i:=start;i<=end;i++{
        dp=max(dp2+nums[i],dp1)
        dp2=dp1
        dp1=dp
    }
    
    return dp
}


func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}