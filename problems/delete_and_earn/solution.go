func deleteAndEarn(nums []int) int {
    var points [10001]int
    
    for i:=0;i<len(nums);i++{
        points[nums[i]]+=nums[i]
    }
    
    prev,cur:=0,0
    for i:=0;i<10000;i++{
        temp:=cur
        cur=max(prev+points[i],cur)
        prev=temp
    }
    
    return cur
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}