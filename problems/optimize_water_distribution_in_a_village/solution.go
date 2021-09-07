func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
    var edges [][]int
    parent:=make([]int,n+1)
    
    for i:=0;i<=n;i++{
        parent[i]=i
    }
    
    for i:=0;i<n;i++{
        edges=append(edges,[]int{0,i+1,wells[i]})
    }
    
    for i:=0;i<len(pipes);i++{
        edges=append(edges,pipes[i])
    }
    
    sort.Slice(edges,func(a,b int) bool{
        return edges[a][2]<edges[b][2]
    })
    
    result:=0
    for i:=0;i<len(edges);i++{
        if union(parent,edges[i][0],edges[i][1]){
            result+=edges[i][2]
        }
    }
    
    return result
}


func find(parent []int,i int) int{
    if parent[i]!=i{
        parent[i]=find(parent,parent[i])
    }
    
    return parent[i]
}

func union(parent []int,x,y int) bool{
    px,py:=find(parent,x),find(parent,y)
    if px!=py{
        parent[px]=py
        return true
    }
    
    return false
}