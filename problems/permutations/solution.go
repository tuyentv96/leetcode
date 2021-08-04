func permute(nums []int) [][]int {
    var result [][]int
    var cur []int
    used:=make(map[int]bool)
    backtrack(nums,&result,cur,used)
    return result
}


func backtrack(nums []int,result *[][]int,cur []int,used map[int]bool){
    if len(cur)==len(nums){
        temp := make([]int, len(nums))
        copy(temp,cur)
        *result=append(*result,temp)
        return
    }
    
    for i:=0;i<len(nums);i++{
        if used[nums[i]]{
            continue
        }
        
        used[nums[i]]=true
        cur=append(cur,nums[i])
        backtrack(nums,result,cur,used)
        cur=cur[:len(cur)-1]
        used[nums[i]]=false
    }
}