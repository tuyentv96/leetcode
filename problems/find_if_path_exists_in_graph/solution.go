func validPath(n int, edges [][]int, start int, end int) bool {
    graph:=make([][]int,n)
    
    for i:=0;i<len(edges);i++{
        a,b:=edges[i][0],edges[i][1]
        graph[a]=append(graph[a],b)
        graph[b]=append(graph[b],a)
    }
    
    visited:=make([]bool,n)
    dfs(graph,visited,start)
    if visited[end]{
        return true
    }
    
    return false
}

func dfs(graph [][]int,visited []bool,i int){
    if visited[i]{
        return
    }
    
    visited[i]=true
    for j:=0;j<len(graph[i]);j++{
        dfs(graph,visited,graph[i][j])
    }
    
    return
}