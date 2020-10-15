/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
   return dfs(p,q)
}

func dfs(p,q *TreeNode) bool{
    if p==nil && q==nil{
        return true
    }
    
    if p==nil || q==nil{
        return false
    }
    
    if p.Val!=q.Val{
        return false
    }
    
    if !dfs(p.Left,q.Left) || !dfs(p.Right,q.Right){
        return false
    }
    
    return true
}