func twoSum(numbers []int, target int) []int {
    var l,r = 0,len(numbers)-1
    var res []int
    for l<r {
        if numbers[l]+numbers[r]==target{
            res=append(res,l+1,r+1)
            break
        }
        
        if numbers[l]+numbers[r]<target{
            l++
        }else{
            r--
        }
    }
    
    return res
}