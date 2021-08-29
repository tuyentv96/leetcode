func minSubArrayLen(target int, nums []int) int {
    const maxInt = 1<<32 - 1
    n:=len(nums)
    start:=0
    sum:=0
    res:=maxInt
    hasValue:=false
    
    for end:=0;end<n;end++{
        sum+=nums[end]
        
        for sum>=target && start<=end{
            res=min(res,end-start+1)
            sum-=nums[start]
            start++
            hasValue=true
        }
    }
    
    if !hasValue{
        return 0
    }
    
    return res
}

func min(a,b int) int{
    if a<b{
        return a
    }
    
    return b
}