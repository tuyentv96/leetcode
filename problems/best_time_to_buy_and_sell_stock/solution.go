func maxProfit(prices []int) int {
    if len(prices)<=1{
        return 0
    }
    
    var res int
    min:=prices[0]
    for i:=0;i<len(prices);i++{
        diff:=prices[i]-min
        if diff>res{
            res=diff
        }
        
        if min>prices[i]{
            min=prices[i]
        }
    }
    
    return res
}