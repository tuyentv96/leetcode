func mincostTickets(days []int, costs []int) int {
    dayMap:=make(map[int]struct{})
    for _,day:=range days{
        dayMap[day]=struct{}{}
    }
    
    var dp [366]int
    for i:=1;i<=365;i++{
        dp[i]=dp[i-1]
        if _,ok:=dayMap[i];ok{
            dp[i]=min(dp[i-1]+costs[0], min(dp[max(0,i-7)]+costs[1], dp[max(0,i-30)]+costs[2]))
        }
    }
    
    return dp[365]
}


func min(a,b int) int{
    if a<b{
        return a
    }
    
    return b
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}