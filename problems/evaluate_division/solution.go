type Node struct{
    ratio float64
    s string
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    parent:=make(map[string]Node)
    for i:=0;i<len(equations);i++{
        dividend:=equations[i][0]
        divisor:=equations[i][1]
        ratio:=values[i]
        union(parent,dividend,divisor,ratio)
    }
    
    result:=make([]float64,len(queries))
    for i:=0;i<len(queries);i++{
        dividend,devisor:=queries[i][0],queries[i][1]
        
        pdividend,ok1:=parent[dividend]
        pdevisor,ok2:=parent[devisor]
        
        if !ok1 || !ok2{
            result[i]=-1
        }else{
            pdividend=find(parent,dividend)
            pdevisor=find(parent,devisor)
            if pdividend.s != pdevisor.s{
                result[i]=-1
            }else{
                result[i]=pdividend.ratio/pdevisor.ratio
            }
        }
    }
    
    
    return result
}


func find(parent map[string]Node,s string) Node{
    if _,ok:=parent[s];!ok{
        parent[s]=Node{s:s,ratio:1}
    }
    
    if parent[s].s!=s{
        newParent:=find(parent,parent[s].s)
        parent[s]=Node{s:newParent.s,ratio:newParent.ratio*parent[s].ratio}
    }
    
    return parent[s]
}

func union(parent map[string]Node,dividend,divisor string,ratio float64) bool{
    pdividend,pdivisor:=find(parent,dividend),find(parent,divisor)
    if pdividend.s!=pdivisor.s{
    parent[pdividend.s]=Node{s:pdivisor.s,ratio:pdivisor.ratio*ratio/pdividend.ratio}
        return true
    }
    
    return false
}