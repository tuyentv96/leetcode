func validTree(n int, edges [][]int) bool {
    if len(edges)!=n-1{
        return false
    }
    
    parent:=make([]int,n)
    
    for i:=0;i<n;i++{
        parent[i]=i
    }
    
    for i:=0;i<len(edges);i++{
        if !union(parent,edges[i][0],edges[i][1]){
            return false
        }
    }
    
    return true
}


func find(parent []int,i int) int{
    if parent[i]!=i{
        parent[i]=find(parent,parent[i])
    }
    
    return parent[i]
}

func union(parent []int,x,y int) bool{
    px:=find(parent,parent[x])
    py:=find(parent,parent[y])
    if px!=py{
        parent[px]=py
        return true
    }
    
    return false
}