func allPathsSourceTarget(graph [][]int) [][]int {  
    var result [][]int
    var tmp []int
    n:=len(graph)
    
    tmp=append(tmp,0)
    dfs(graph,&result,tmp,0,n-1)
    
    return result
}

func dfs(graph [][]int,result *[][]int,cur []int,src,dest int) {
    if src==dest{
        tmp:=make([]int,len(cur))
        copy(tmp,cur)
        *result=append(*result,tmp)
        return
    }

    for i:=0;i<len(graph[src]);i++{
        cur=append(cur,graph[src][i])
        dfs(graph,result,cur,graph[src][i],dest)
        cur=cur[:len(cur)-1]
    }
        
    return
}