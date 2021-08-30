func runningSum(nums []int) []int {
    result:=make([]int,0,len(nums))
    sum:=0
    for i:=0;i<len(nums);i++{
        sum+=nums[i]
        result=append(result,sum)
    }
    
    return result
}