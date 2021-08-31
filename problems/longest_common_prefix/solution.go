func longestCommonPrefix(strs []string) string {
    sort.Slice(strs,func (a,b int) bool{
        return strs[a]<strs[b]
    })
    
    first:=strs[0]
    last:=strs[len(strs)-1]
    
    for i:=0;i<len(first);i++{
        if first[i]!=last[i]{
            return first[:i]
        }
    }
    
    return first
}