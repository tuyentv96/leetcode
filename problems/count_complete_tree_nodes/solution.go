/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root==nil{
        return 0
    }
    
    ld:=leftDepth(root.Left)
    rd:=leftDepth(root.Right)
    if ld==rd{
        return 1<<ld + countNodes(root.Right)
    }
    
    return 1<<rd + countNodes(root.Left)
}

func leftDepth(root *TreeNode) int{
    if root==nil{
        return 0
    }
    
    return 1 + leftDepth(root.Left)
}