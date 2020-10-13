func findLengthOfLCIS(nums []int) int {
    var res int
    for r:=0;r<len(nums);r++{
        l:=r
        for r<len(nums)-1 && nums[r]<nums[r+1]{
            r++
        }
        
        if res<(r-l+1){
            res=r-l+1
        }
    }
    
    return res
}