const intMax = 1<<32

func minCostConnectPoints(points [][]int) int {
    n:=len(points)
    keys:=make([]int,n)
    visited:=make([]bool,n)
    for i:=0;i<n;i++{
        keys[i]=intMax
    }
    
    keys[0]=0
    visited[0]=true
    cost:=0
    for i:=0;i<n;i++{
        idx:=minKey(keys,visited)
        visited[idx]=true
        cost+=keys[idx]
        
        for j:=0;j<n;j++{
            v:=abs(points[idx][0]-points[j][0])+abs(points[idx][1]-points[j][1])
            if !visited[j] && (v < keys[j]){
                keys[j]=v
            }
        }
    }
    
    return cost
}

func minKey(keys []int,visited []bool) int{
    min,idx:=intMax,0
    
    for i:=0;i<len(keys);i++{
        if !visited[i] && keys[i]<min{
            min=keys[i]
            idx=i
        }
    }
    
    return idx
}

func abs(a int) int{
    if a<0{
        return -a
    }
    
    return a
}