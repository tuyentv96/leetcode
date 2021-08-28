/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    helper(root)
    return root
}


func helper(root *TreeNode) {
    if root==nil{
        return
    }
    
    root.Left,root.Right=root.Right,root.Left
    helper(root.Left)
    helper(root.Right)
}