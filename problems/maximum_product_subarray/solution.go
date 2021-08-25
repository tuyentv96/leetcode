func maxProduct(nums []int) int {
    if len(nums)==0{
        return 0
    }
    
    max_so_far,min_so_far:=nums[0],nums[0]
    result:=max_so_far
    
    for i:=1;i<len(nums);i++{
        cur:=nums[i]
        tmp:=max(cur,max(max_so_far*cur,min_so_far*cur))
        min_so_far=min(cur,min(max_so_far*cur,min_so_far*cur))
        
        max_so_far=tmp
        result=max(result,max_so_far)
    }
    
    return result
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}

func min(a,b int) int{
    if a<b{
        return a
    }
    
    return b
}