func findOrder(numCourses int, prerequisites [][]int) []int {
    graph:=make(map[int][]int)
    
    for i:=0;i<len(prerequisites);i++{
        a:=prerequisites[i][0]
        b:=prerequisites[i][1]
        graph[b]=append(graph[b],a)
    }
    
    result:=make([]int,0,numCourses)
    visited:=make([]bool,numCourses)
    memo:=make([]bool,numCourses)
    
    for i:=0;i<numCourses;i++{
        if !dfs(graph,memo,visited,i,&result){
            return nil
        }
    }
    
    return result
}

func dfs(graph map[int][]int,memo,visited []bool,i int,result *[]int) bool{
    if visited[i]{
        return memo[i]
    }
    
    visited[i]=true
    
    for j:=0;j<len(graph[i]);j++{
        if !dfs(graph,memo,visited,graph[i][j],result){
            memo[i]=false
            return false
        }
    }
    
    memo[i]=true
    *result=append([]int{i},*result...)
    return true
}