func earliestAcq(logs [][]int, n int) int {
    parent:=make([]int,n)
    
    for i:=0;i<n;i++{
        parent[i]=i
    }
    
    sort.Slice(logs,func(i,j int) bool{
        return logs[i][0]<logs[j][0]
    })
    
    count:=n
    for i:=0;i<len(logs);i++{
        ts,src,dest:=logs[i][0],logs[i][1],logs[i][2]
        if union(parent,src,dest){
            count--
        }
        
        if count==1{
            return ts
        }
    }
    
    return -1
}

func find(parent []int,i int) int{
    if parent[i]!=i{
        parent[i]=find(parent,parent[i])
    }
    
    return parent[i]
}

func union(parent []int,x,y int) bool{
    px:=find(parent,x)
    py:=find(parent,y)
    
    if px!=py{
        parent[px]=py
        return true
    }
    
    return false
}