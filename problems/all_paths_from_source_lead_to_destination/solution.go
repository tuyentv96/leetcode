type Color int

const (
    ColorWhite Color = iota
    ColorGray
    ColorBlack
)

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
    graph:=make([][]int,n)
    
    for i:=0;i<len(edges);i++{
        src,dest:=edges[i][0],edges[i][1]
        graph[src]=append(graph[src],dest)
    }
    
    visited:=make([]bool,n)
    colors:=make([]Color,n)
    return dfs(graph,visited,colors,source,destination)
}

func dfs(graph [][]int,visited []bool,colors []Color,i int,dest int) bool{
    if colors[i]!=ColorWhite{
        return colors[i]==ColorBlack
    }
    
    if len(graph[i])==0{
        return i==dest
    }
    
    colors[i]=ColorGray
    
    // visited[i]=true
    for j:=0;j<len(graph[i]);j++{
        if !dfs(graph,visited,colors,graph[i][j],dest){
            return false
        }
    }
    
    colors[i]=ColorBlack
    // visited[i]=false
    return true
}