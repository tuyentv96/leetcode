func generateParenthesis(n int) []string {
	var arr []string
	backtrack(&arr, "", 0, 0, n)
	return arr
}

func backtrack(arr *[]string, cur string, open, close, n int) {
    if len(cur)==n*2{
        *arr=append(*arr,cur)
        return
    }
    
    if open<n{
        backtrack(arr,cur+"(",open+1,close,n)
    }
    
    if close<open{
        backtrack(arr,cur+")",open,close+1,n)
    }
}
