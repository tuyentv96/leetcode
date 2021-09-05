func countComponents(n int, edges [][]int) int {
    parent:=make([]int,n)

    for i:=0;i<n;i++{
        parent[i]=i
    }
    
    for i:=0;i<len(edges);i++{
        union(parent,edges[i][0],edges[i][1])
    }
    
    count:=0
    for i:=0;i<n;i++{
        if parent[i]==i{
            count++
        }
    }
    
    return count
}

func find(parent []int,x int) int{
    if parent[x]!=x{
        parent[x]=find(parent,parent[x])
    }
    
    return parent[x]
}

func union(parent []int,x,y int) {
    px:=find(parent,x)
    py:=find(parent,y)
    
    if px!=py{
        parent[px]=py
    }
}