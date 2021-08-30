func findMaxConsecutiveOnes(nums []int) int {
    count,res:=0,0
    
    for i:=0;i<len(nums);i++{
        if nums[i]==1{
            count++
        }else{
            res=max(res,count)
            count=0
        }
        
    }
    
    res=max(res,count)
    return res
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}