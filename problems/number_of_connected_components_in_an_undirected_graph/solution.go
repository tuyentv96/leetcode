func countComponents(n int, edges [][]int) int {
    graph:=make([][]int,n)
    
    for i:=0;i<len(edges);i++{
        graph[edges[i][0]]=append(graph[edges[i][0]],edges[i][1])
        graph[edges[i][1]]=append(graph[edges[i][1]],edges[i][0])
    }
    
    visited:=make([]bool,n)
    memo:=make([]int,n)
    count:=0
    
    for i:=0;i<n;i++{
        if !visited[i]{
            count++
            dfs(graph,visited,memo,i)
        }
    }
    
    return count
}

func dfs(graph [][]int,visited []bool,memo []int, node int){
    if visited[node]{
        return
    }
    
    visited[node]=true
    for i:=0;i<len(graph[node]);i++{
        dfs(graph,visited,memo,graph[node][i])
    }
    
    return
}