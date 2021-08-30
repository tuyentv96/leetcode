func subarraysDivByK(nums []int, k int) int {
    sum:=0
    m:=make(map[int]int)
    m[0]=1
    res:=0
    
    for i:=0;i<len(nums);i++{
        sum = (sum + nums[i]) % k
        
        if sum<0{
            sum+=k
        }
        
        res+=m[sum]
        m[sum]++
    }
    
    return res
}
