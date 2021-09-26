func minCostConnectPoints(points [][]int) int {
    edges:=make([][]int,0)
    n:=len(points)
    
    for i:=0;i<n;i++{
        for j:=i+1;j<n;j++{
            edges=append(edges,[]int{i,j,abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])})
        }
    }
    
    sort.Slice(edges,func (a,b int)bool{
        return edges[a][2]<edges[b][2]
    })
    
    parent:=make([]int,n)
    
    for i:=0;i<n;i++{
        parent[i]=i    
    }
    
    result:=0
    for i:=0;i<len(edges);i++{
        x,y,cost:=edges[i][0],edges[i][1],edges[i][2]
        if union(parent,x,y){
            result+=cost
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

func abs(a int) int{
    if a<0{
        return -a
    }
    
    return a
}