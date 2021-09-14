var directions = [][]int{{-1,-1},{-1,0},{-1,1},{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1}}


type Queue struct{
    arr [][]int
}

func NewQueue() *Queue{
    return &Queue{arr:make([][]int,0)}
}

func (q *Queue) Enqueue(e []int) {
    q.arr=append(q.arr,e)
}

func (q *Queue) Dequeue() []int{
    e:=q.arr[0]
    q.arr=q.arr[1:]
    return e
}

func (q *Queue) IsEmpty() bool{
    return len(q.arr)==0
}

func shortestPathBinaryMatrix(grid [][]int) int {
    if len(grid)==0 || grid[0][0]!=0{
        return -1
    }
    
    queue:=NewQueue()
    n:=len(grid)
    grid[0][0]=1
    queue.Enqueue([]int{0,0})
    
    for !queue.IsEmpty(){
        e:=queue.Dequeue()
        row,col:=e[0],e[1]
        
        distance:=grid[row][col]
        if row==n-1 && col==n-1{
            return distance
        }
        
        neighbors:=getNeighbors(grid,e[0],e[1])
        
        for i:=0;i<len(neighbors);i++{
            nrow,ncol:=neighbors[i][0],neighbors[i][1]
            queue.Enqueue([]int{nrow,ncol})
            grid[nrow][ncol]=distance+1
        }
    }
    
    return -1
}

func getNeighbors(grid [][]int,i,j int) [][]int{
    result:=make([][]int,0)
    n:=len(grid)
    
    for k:=0;k<len(directions);k++{
        row,col:=i+directions[k][0],j+directions[k][1]
        if row < 0 || col <0 || row>=n || col>=n || grid[row][col]!=0{
            continue
        }
        
        result=append(result,[]int{row,col})
    }
    
    return result
}