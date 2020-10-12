/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root==nil{
        return true
    }
    
    lh:=height(root.Left)
    rh:=height(root.Right)
    
    var i int
    if lh>rh{
     i=lh-rh   
    }else{
        i=rh-lh
    }
    
    if i<=1 && isBalanced(root.Left) && isBalanced(root.Right){
        return true
    }
    
    return false
}


func height(root *TreeNode) int{
    if root==nil{
        return 0
    }
    
    lh:=height(root.Left)
    rh:=height(root.Right)
    if lh > rh{
        return 1 + lh
    }
    
    return 1 + rh
}