/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited:=make(map[*Node]*Node)
    
    return dfs(node,visited)
}

func dfs(node *Node,visited map[*Node]*Node) *Node{
    if node==nil{
        return nil
    }
    
    if visited[node] != nil{
        return visited[node]
    }
    
    cloneNode:=&Node{
        Val: node.Val,
    }
    
    visited[node]=cloneNode
    
    for i:=0;i<len(node.Neighbors);i++{
        cloneNode.Neighbors=append(cloneNode.Neighbors,dfs(node.Neighbors[i],visited))
    }
    
    return cloneNode
}