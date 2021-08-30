func pivotIndex(nums []int) int {
    n:=len(nums)
    sum,leftSum:=0,0
    
    for i:=0;i<n;i++{
        sum+=nums[i]
    }
    
    for i:=0;i<n;i++{
        if leftSum==(sum-leftSum-nums[i]){
            return i
        }
        
        leftSum+=nums[i]
    }
    
    return -1
}