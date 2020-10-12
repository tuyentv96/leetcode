/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    m:=make(map[int]int)
    var maxCount int
    traversal(root,m,&maxCount)
    var res []int
    for k,v:=range m{
        if v==maxCount{
            res=append(res,k)
        }
    }
    
    return res
}

func traversal(root *TreeNode,m map[int]int,maxCount *int) {
    if root==nil{
        return
    }
    
    if root.Left!=nil{
         traversal(root.Left,m,maxCount)
    }
    
    m[root.Val]+=1
    if *maxCount<m[root.Val]{
        *maxCount=m[root.Val]
    }

    if root.Right!=nil{
         traversal(root.Right,m,maxCount)
    }
}