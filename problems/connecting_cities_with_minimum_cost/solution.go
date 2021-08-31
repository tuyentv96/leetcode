func minimumCost(n int, connections [][]int) int {
    parents:=make([]int,n+1)
    result:=0
    
    for i:=0;i<=n;i++{
        parents[i]=i
    }
    
    sort.Slice(connections,func (x,y int) bool{
        return connections[x][2]<connections[y][2]
    })
    
    uf:=UnionFind{
        count: n,
        parents: parents,
    }
    
    for i:=0;i<len(connections);i++{
        x,y,cost:=connections[i][0],connections[i][1],connections[i][2]
        if uf.union(x,y){
            result+=cost
        }
    }
    
    if uf.count!=1{
        result=-1
    }
    
    return result
}

type UnionFind struct{
    count int
    parents []int
}

func (uf *UnionFind) find(x int) int{
    if uf.parents[x]!=x{
        uf.parents[x]=uf.find(uf.parents[x])
    }
    
    return uf.parents[x]
}

func (uf *UnionFind) union(x,y int) bool{
    px,py:=uf.find(x),uf.find(y)
    if px!=py{
        uf.parents[px]=py
        uf.count--
        return true
    }
    
    return false
}