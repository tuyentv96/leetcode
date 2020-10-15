/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
   return isMirror(root,root)
}

func isMirror(l,r *TreeNode) bool{
    if l==nil && r==nil{
        return true
    }
    
    if l!=nil && r!=nil && l.Val == r.Val{
        return isMirror(l.Left,r.Right) && isMirror(l.Right,r.Left)
    }
    
    return false
}