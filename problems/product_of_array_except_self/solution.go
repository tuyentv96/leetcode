func productExceptSelf(nums []int) []int {
    length:=len(nums)
    
    result:=make([]int,length)
    result[0]=1
    
    for i:=1;i<length;i++{
        result[i]=result[i-1]*nums[i-1]
    }
    
    right:=1
    for i:=length-1;i>=0;i--{
        result[i]=result[i]*right
        right*=nums[i]
    }
    
    return result
}