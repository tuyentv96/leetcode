/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    res:=make([][]int,0)
    traversal(root,&res,0)
    return res
}

func traversal(root *TreeNode,res *[][]int,level int){
    if root==nil{
        return
    }
    
    if len(*res)<=level{
        *res=append(*res,make([]int,0))
    }
    
    (*res)[level]=append((*res)[level],root.Val)
    traversal(root.Left,res,level+1)
    traversal(root.Right,res,level+1)
}
