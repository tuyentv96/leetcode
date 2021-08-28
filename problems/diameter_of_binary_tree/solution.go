/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var result int
    helper(root,&result)
    return result
}


func helper(root *TreeNode,result *int) int{
    if root==nil{
        return 0
    }
    
    left,right:=helper(root.Left,result),helper(root.Right,result)

    *result=max(*result,left+right)

    return max(left,right)+1
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}