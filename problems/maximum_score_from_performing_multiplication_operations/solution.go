// func maximumScore(nums []int, multipliers []int) int {
//     memo:=make([][]int,len(multipliers))
//     for i:=range memo{
//         memo[i]=make([]int,len(multipliers))
//     }
    
//     return helper(nums,multipliers,memo,0,0)
// }

// func helper(nums []int, multipliers []int,memo [][]int,i,left int) int{
//     if i==len(multipliers){
//         return 0
//     }
    
//     mul:=multipliers[i]
//     right:=len(nums)-1 - (i-left)
    
//     if memo[i][left]==0{
//         memo[i][left]=max(nums[left]*mul + helper(nums,multipliers,memo,i+1,left+1),nums[right]*mul +helper(nums,multipliers,memo,i+1,left))
//     }
    
//     return memo[i][left]
// }

func maximumScore(nums []int, multipliers []int) int {
    dp:=make([][]int,len(multipliers)+1)
    for i:=range dp{
        dp[i]=make([]int,len(multipliers)+1)
    }
    
    for i:=len(multipliers)-1;i>=0;i--{
        for left:=i;left>=0;left--{
            mul:=multipliers[i]
            right:=len(nums)-1 - (i-left)
            dp[i][left]=max(nums[left]*mul+dp[i+1][left+1],nums[right]*mul+dp[i+1][left])
        }
    }
    
    return dp[0][0]
}


func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}