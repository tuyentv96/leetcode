func validPath(n int, edges [][]int, start int, end int) bool {
    graph:=make([][]int,n)
    
    for i:=0;i<len(edges);i++{
        a,b:=edges[i][0],edges[i][1]
        graph[a]=append(graph[a],b)
        graph[b]=append(graph[b],a)
    }
    
    visited:=make([]bool,n)
    queue:=NewQueue()
    queue.Enqueue(start)
    
    return bfs(graph,visited,queue,end)
}

func bfs(graph [][]int,visited []bool,queue Queue,end int) bool{
    for !queue.IsEmpty(){
        i:=queue.Dequeue()
        
        if i==end{
            return true
        }
        
        visited[i]=true
        
        for j:=0;j<len(graph[i]);j++{
            k:=graph[i][j]
            if !visited[k]{
                queue.Enqueue(k)
            }
        }
    }
    
    return false
}

type Queue struct{
    arr []int
}

func NewQueue() Queue{
    return Queue{
        arr: make([]int,0),
    }
}

func (q *Queue) IsEmpty()bool{
    return len(q.arr)==0
}


func (q *Queue) Enqueue(i int) {
    q.arr=append(q.arr,i)
}

func (q *Queue) Dequeue() int{
    i:=q.arr[0]
    q.arr=q.arr[1:]
    return i
}