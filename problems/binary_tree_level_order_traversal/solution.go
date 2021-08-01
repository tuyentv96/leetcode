/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root==nil{
        return [][]int{}
    }
    
    depth:=maxDepth(root)
    // var result [][]int
    result:=make([][]int,depth)
    helper(root,&result,0)
    return result
}

func helper(root *TreeNode,result *[][]int,level int){
    if root==nil{
        return
    }
    
    // if len(*result)<=level{
    //     *result=append(*result,[]int{})
    // }
    
    (*result)[level]=append((*result)[level],root.Val)
    helper(root.Left,result,level+1)
    helper(root.Right,result,level+1)
}

func maxDepth(root *TreeNode) int{
    if root==nil{
        return 0
    }
    
    return max(maxDepth(root.Left),maxDepth(root.Right)) + 1
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}