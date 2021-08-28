func characterReplacement(s string, k int) int {
    var charCount [26]int
    var maxCharCount int
    start:=0
    var result int

    for end:=0;end<len(s);end++{
        c:=s[end]-'A'
        charCount[c]++
        maxCharCount=max(maxCharCount,charCount[c])
        
        if end-start+1-maxCharCount>k{
            charCount[s[start]-'A']--
            start++
        }
        
        result=max(result,end-start+1)
    }
    
    return result
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}