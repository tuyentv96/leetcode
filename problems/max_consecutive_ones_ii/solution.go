func findMaxConsecutiveOnes(nums []int) int {
    start,res:=0,0
    count,countZero:=0,0
    
    for end:=0;end<len(nums);end++{
        if nums[end]==0{
            countZero++
        }
        
        count++
        
        for countZero>1 && start<=end{
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