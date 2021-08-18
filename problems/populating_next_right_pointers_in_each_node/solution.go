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
    dfs(root,nil)
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