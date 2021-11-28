/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1==nil && root2==nil{
        return nil
    }
    

    
    result:=&TreeNode{}    
    var root1L,root1R,root2L,root2R *TreeNode

    if root1!=nil{
        result.Val+=root1.Val
        root1L=root1.Left
        root1R=root1.Right
    }
    
    if root2!=nil{
        result.Val+=root2.Val
        root2L=root2.Left
        root2R=root2.Right
    }
    
    result.Left=mergeTrees(root1L,root2L)
    result.Right=mergeTrees(root1R,root2R)
    
    return result
}
