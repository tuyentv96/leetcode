func longestOnes(nums []int, k int) int {
    countZero:=0
    start:=0
    res:=0
    
    for end:=0;end<len(nums);end++{
        if nums[end]==0{
            countZero++
        }
        
        for countZero>k{
            if nums[start]==0{
                countZero--
            }
            
            start++
        }
        
        res=max(res,end-start+1)
    }
    
    return res
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}