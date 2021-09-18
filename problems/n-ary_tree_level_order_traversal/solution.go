/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Queue struct{
    list *list.List
}

func NewQueue() *Queue{
    return &Queue{
        list: list.New(),
    }
}

func (q *Queue) Enqueue(v *Node){
    q.list.PushBack(v)
}

func (q *Queue) Dequeue() *Node{
    e:=q.list.Front()
    q.list.Remove(e)
    return e.Value.(*Node)
}

func (q *Queue) Len() int{
    return q.list.Len()
}

func (q *Queue) IsEmpty() bool{
    return q.list.Len()==0
}

func levelOrder(root *Node) [][]int {
    if root==nil{
        return nil
    }
    
    queue:=NewQueue()
    result:=make([][]int,0)
    queue.Enqueue(root)
    
    for !queue.IsEmpty(){
        size:=queue.Len()
        arr:=make([]int,0,size)
        
        for i:=0;i<size;i++{
            e:=queue.Dequeue()
            arr=append(arr,e.Val)
            
            for j:=0;j<len(e.Children);j++{
                queue.Enqueue(e.Children[j])
            }
        }
        
        result=append(result,arr)
    }
    
    return result
}