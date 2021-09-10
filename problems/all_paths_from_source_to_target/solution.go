type Queue struct{
    arr [][]int
}

func NewQueue() Queue{
    return Queue{
        arr: make([][]int,0),
    }
}

func (q *Queue) IsEmpty()bool{
    return len(q.arr)==0
}


func (q *Queue) Enqueue(i []int) {
    q.arr=append(q.arr,i)
}

func (q *Queue) Dequeue() []int{
    i:=q.arr[0]
    q.arr=q.arr[1:]
    return i
}

func allPathsSourceTarget(graph [][]int) [][]int {  
    var result [][]int
    n:=len(graph)
    path:=[]int{0}
    queue:=NewQueue()
    queue.Enqueue(path)
    
    bfs(graph,&result,queue,n-1)
    
    return result
}

func clone(from []int) []int{
    tmp:=make([]int,len(from))
    copy(tmp,from)
    return tmp
}

func bfs(graph [][]int,result *[][]int,queue Queue,dest int) {
    for !queue.IsEmpty(){
        path:=queue.Dequeue()
        last:=path[len(path)-1]
        
        if last==dest{
            *result=append(*result,clone(path))
        }
        
        for i:=0;i<len(graph[last]);i++{
            newPath:=clone(path)
            newPath=append(newPath,graph[last][i])
            queue.Enqueue(newPath)
        }
    }

    return
}