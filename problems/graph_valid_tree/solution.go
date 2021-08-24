func validTree(n int, edges [][]int) bool {
    graph:=make([][]int,n)
    
    for i:=0;i<len(edges);i++{
        a,b:=edges[i][0],edges[i][1]
        graph[a]=append(graph[a],b)
        graph[b]=append(graph[b],a)
    }
    
    visited:=make([]bool,n)
    
    // check graph has cycle
    if isCycle(graph,visited,0,-1){
        return false
    }

    
    // check graph is connected
    for i:=0;i<n;i++{
        if !visited[i]{
            return false
        }
    }
    
    return true
}


func isCycle(graph [][]int,visited []bool, node int,parent int) bool{
    if visited[node]{
        return true
    }
    
    visited[node]=true

    for i:=0;i<len(graph[node]);i++{
        v:=graph[node][i]
        if v!=parent && isCycle(graph,visited,v,node){
            return true
        }
    }
    
    return false
}