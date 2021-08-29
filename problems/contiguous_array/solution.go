func findMaxLength(nums []int) int {
    m:=make(map[int]int)
    n:=len(nums)
    count:=0
    res:=0
    m[0]=-1
    
    for i:=0;i<n;i++{
        if nums[i]==0{
            count--
        }else{
            count++
        }
        
        if index,ok:=m[count];ok{
            res=max(res,i-index)
        }else{
            m[count]=i
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