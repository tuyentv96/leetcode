/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    var result []string
    var cur string
    helper(root,&result,cur)
    return result
}


func helper(root *TreeNode,result *[]string,cur string){
    if root==nil{
        return
    }

    
    if cur==""{
        cur+= fmt.Sprintf("%d",root.Val)    
    }else{
        cur+= fmt.Sprintf("->%d",root.Val)
    }
    
    if root.Left==nil && root.Right==nil{
        *result=append(*result, cur)
    }
    
    helper(root.Left,result,cur)
    helper(root.Right,result,cur)
}