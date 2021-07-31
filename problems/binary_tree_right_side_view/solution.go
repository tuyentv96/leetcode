/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var result []int
    dfs(root,&result,1)
    return result
}

func dfs(root *TreeNode, result *[]int,level int){
    if root==nil{
        return
    }
    
    if len(*result)<level{
        *result=append(*result, root.Val)
    }
    dfs(root.Right,result,level+1)
    dfs(root.Left,result,level+1)
}