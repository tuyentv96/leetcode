func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
    for i:=0;i<len(nums);i++{
        j:=target-nums[i]
        if index,ok:=m[j];ok{
            return []int{i,index}
        }
        
        m[nums[i]]=i
    }
    
    return nil
}