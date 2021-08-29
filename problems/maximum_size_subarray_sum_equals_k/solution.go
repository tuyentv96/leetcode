func maxSubArrayLen(nums []int, k int) int {
    sum:=0
    m:=make(map[int]int)
    res:=0
    // m[0] = -1
    
    for i:=0;i<len(nums);i++{
        sum+=nums[i]
        
        if sum==k{
            res=i+1
        }
        
        if index,ok:=m[sum-k];ok{
            res=max(res,i-index)
        }
        
        if _,ok:=m[sum];!ok{
            m[sum]=i
        }
    }
    
    return res
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}