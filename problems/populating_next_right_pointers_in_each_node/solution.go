/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    cur:=root
    var prev,head *Node
    for cur!=nil{
        
        for cur!=nil{
            if cur.Left!=nil{
                if prev!=nil{
                    prev.Next=cur.Left
                }else{
                    head=cur.Left
                }
                
                prev=cur.Left
            }
            
            if cur.Right!=nil{
                if prev!=nil{
                    prev.Next=cur.Right
                }else{
                    head=cur.Right
                }
                
                prev=cur.Right
            }
            
            cur=cur.Next
        }
        
        cur=head
        prev=nil
        head=nil
    }
    
    return root
}

func dfs(cur *Node, next *Node){
    if cur==nil{
        return
    }
    
    cur.Next=next
    dfs(cur.Left,cur.Right)
    
    var rightNext *Node
    if cur.Next!=nil{
        rightNext=cur.Next.Left
    }
    
    dfs(cur.Right,rightNext)
}