/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    h:=height(root)
    res:=make([][]int,h)
    traversal(root,&res,0,h)
    return res
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}

func height(root *TreeNode) int{
    if root==nil{
        return 0
    }
    
    return 1 + max(height(root.Left),height(root.Right))
}

func traversal(root *TreeNode,res *[][]int,level int,h int){
    if root==nil{
        return
    }
    
    traversal(root.Left,res,level+1,h)
    traversal(root.Right,res,level+1,h)
    
    (*res)[h-level-1]=append((*res)[h-level-1],root.Val)
}