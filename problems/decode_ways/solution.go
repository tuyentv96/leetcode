func numDecodings(s string) int {
    dp:=make([]int,len(s)+1)
    
    dp[0]=1
    
    if s[0]!='0'{
        dp[1]=1
    }
    
    for i:=2;i<len(dp);i++{
        oneDigit:=s[i-1]
        twoDigit:=s[i-2:i]
        
        if oneDigit!='0'{
            dp[i]=dp[i-1]
        }
        
        if twoDigit>="10" && twoDigit<="26"{
            dp[i]+=dp[i-2]
        }
    }
    
    return dp[len(s)]
}