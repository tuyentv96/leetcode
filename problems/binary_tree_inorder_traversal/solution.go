/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var result []int
    dfs(root,&result)
    return result
}

func dfs(root *TreeNode, result *[]int){
    if root==nil{
        return
    }
    
    dfs(root.Left,result)
    *result=append(*result,root.Val)
    dfs(root.Right,result)
}