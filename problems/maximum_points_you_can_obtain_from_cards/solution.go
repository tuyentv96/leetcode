func maxScore(cardPoints []int, k int) int {
    n:=len(cardPoints)
    sum:=0
    
    for i:=0;i<k;i++{
        sum+=cardPoints[i]
    }
    
    newSum:=sum
    
    for i:=0;i<k;i++{
        newSum-=cardPoints[k-1-i]
        newSum+=cardPoints[n-1-i]
        
        sum=max(sum,newSum)
    }
    
    return sum
}


func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}