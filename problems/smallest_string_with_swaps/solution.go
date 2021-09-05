func smallestStringWithSwaps(s string, pairs [][]int) string {
    n:=len(s)
    parent:=make([]int,n)
    
    for i:=0;i<n;i++{
        parent[i]=i
    }
    
    for i:=0;i<len(pairs);i++{
        union(parent,pairs[i][0],pairs[i][1])
    }
    
    m:=make(map[int][]int)
    for i:=0;i<n;i++{
        m[find(parent,i)]=append(m[find(parent,i)],i)
    }
    
    result:=make([]byte,n)
    for _,v:=range m{
        tmp:=make([]int,0,len(v))
        for i:=0;i<len(v);i++{
            tmp=append(tmp,v[i])
        }
        
        sort.Slice(tmp,func(i,j int)bool{
            return s[tmp[i]] < s[tmp[j]]
        })
        
        for i:=0;i<len(v);i++{
            result[v[i]]=s[tmp[i]]
        }
    }
    
    return string(result)
}


func find(parent []int, i int) int{
    if parent[i]!=i{
        parent[i]=find(parent,parent[i])
    }
    
    return parent[i]
}

func union(parent []int,x,y int){
    px,py:=find(parent,x),find(parent,y)
    
    if px!=py{
        parent[px]=py
    }
}