/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return traversal(root,nil,nil)
}


func traversal(root *TreeNode,lower,upper *int) bool{
    if root==nil{
        return true
    }
    
    if lower!=nil && root.Val <= *lower{
        return false
    }
    
    if upper!=nil && root.Val >= *upper{
        return false
    }

    
    if !traversal(root.Left,lower,&root.Val) || !traversal(root.Right,&root.Val,upper){
        return false
    }
    
    return true
}