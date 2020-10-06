func removeDuplicates(nums []int) int {
    var count int
    for i:=1;i<len(nums);i++{
        if nums[i]!=nums[count]{
            count++
            nums[count], nums[i] = nums[i],nums[count]
        }
    }
    
    return count+1
}