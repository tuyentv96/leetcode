func climbStairs(n int) int {
    if n<1{
        return 0
    }
    
    if n==1{
        return 1
    }
    
    if n==2{
        return 2
    }
    
    var res = 2
    var prev = 1
    for i:=3;i<=n;i++{
        tmp:=res
        res=res+prev
        prev=tmp
    }
    
    return res
}