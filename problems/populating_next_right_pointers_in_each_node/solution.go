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
