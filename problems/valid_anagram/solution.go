func isAnagram(s string, t string) bool {
    var charSet1 [26]byte
    var charSet2 [26]byte
    
    for i:=0;i<len(s);i++{
        charSet1[s[i]-'a']++
    }
    
    for i:=0;i<len(t);i++{
        charSet2[t[i]-'a']++
    }
    
    for i:=0;i<26;i++{
        if charSet1[i]!=charSet2[i]{
            return false
        }
    }
    
    return true
}